package routers

import (
	"dpm/common"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	IN_PATH  = "inpath"
	IN_QUERY = "inquery"
	IN_BODY  = "inbody"
)

func ParseQueryGet(r *http.Request, key string) string {
	vs, err := url.ParseQuery(r.URL.RawQuery)
	if err == nil && len(vs[key]) > 0 {
		return strings.TrimSpace(vs[key][0])
	}
	return ""
}

func HttpHandlerWrap(r Route) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		start := time.Now()
		common.Logger.Infof("%s\t%s\t%s", req.Method, req.RequestURI, r.Name)

		fv := reflect.ValueOf(r.HandlerFunc)
		ft := reflect.TypeOf(r.HandlerFunc)

		//Got the number of In params
		numIn := ft.NumIn()

		in := make([]reflect.Value, numIn)

		//for each the In param
		for i := 0; i < numIn; i++ {
			//Got the In param's Type
			fit := ft.In(i)
			if fit.Kind() == reflect.Interface {
				if reflect.TypeOf(w).Implements(fit) { // is http.ResponseWriter
					in[i] = reflect.ValueOf(w)
					continue
				}
			}
			if fit == reflect.TypeOf(req) { // is http.Request
				in[i] = reflect.ValueOf(req)
				continue
			}

			if fit.Kind() != reflect.Struct { 
				// It'll panic a error when the param is not the struct type
				panic(common.ErrInternalServerS(`the request params mapping support "struct" type only...`))
			}

			in[i] = structEvaluate(fit, req)
			_, err := govalidator.ValidateStruct(in[i].Interface())
			if err != nil {
				panic(common.ErrTrace(err))
			}
		}
		out := fv.Call(in)
		common.Logger.Infof("%s\t%s\t%s\t%s", req.Method, req.RequestURI, r.Name, time.Since(start))
		reponseOut(w, out)
	})
}

func isErr(o interface{}) (bool, error) {
	if err, ok := o.(error); ok {
		return true, err
	}
	return false, nil
}

func reponseOut(w http.ResponseWriter, outs []reflect.Value) {

	if outs == nil || len(outs) == 0 {
		common.JsonResponseMsg(w, http.StatusOK, "ok")
	} else if len(outs) == 1 {
		item := outs[0].Interface()
		if item == nil {
			common.JsonResponseMsg(w, http.StatusOK, "ok")
		} else if yes, err := isErr(item); yes {
			panic(common.ErrTrace(err))
		} else {
			common.JsonResponseOK(w, &item)
		}
	} else {
		outSli := make([]interface{}, len(outs))
		countSli := 0
		for idx, item := range outs {
			itemInf := item.Interface()
			if yes, err := isErr(itemInf); yes {
				panic(common.ErrTrace(err))
			}
			if itemInf == nil {
				continue
			}
			outSli[idx] = itemInf
			countSli++
		}
		switch countSli {
		case 0:
			common.JsonResponseMsg(w, http.StatusOK, "ok")
		case 1:
			common.JsonResponseOK(w, &outSli[0])
		default:
			common.JsonResponseOK(w, &outSli)
		}
	}
}

func structEvaluate(t reflect.Type, req *http.Request) reflect.Value {
	nt := reflect.New(t)
	v := nt.Elem()
	tnf := t.NumField()
	noTagStruct := true
	for i := 0; i < tnf; i++ {
		//迭代每一个field
		tf := t.Field(i)
		fmt.Println(tf.Name)
		tft := tf.Type
		tftk := tft.Kind()
		//获取tag元数据，确定取值来源
		ftag := tf.Tag
		tag, ok := ftag.Lookup("qval")
		if !ok {
			continue
		} else {
			noTagStruct = false
		}

		var tfv reflect.Value

		tagElems := strings.Split(tag, ",")
		elemLen := len(tagElems)
		switch elemLen {
		case 0:
			panic(common.ErrInternalServerf(`
				the "%s" tag's element value in ["%s\t%s\t%s\t%s\t"] at least one. `,
				"qval", IN_PATH, IN_QUERY, IN_BODY, "+", "-"))
		case 1:
			eleFst := tagElems[0]
			switch eleFst {
			case "-":
				continue
			case "+":
				if tftk != reflect.Struct {
					panic(common.ErrInternalServerf(`
						the "%s" tag's element [%s] support struct filed only. `,
						"qval", "+"))
				}
				tfv = structEvaluate(tft, req)
			case IN_BODY:
				if tftk != reflect.Struct {
					panic(common.ErrInternalServerf(`
						the "%s" tag's element [%s] support struct filed only. `,
						"qval", IN_BODY))
				}
				o := reflect.New(tft).Interface()
				if err := common.Unmarshal2Object(req, &o); err!= nil {
					panic(common.ErrBadRequestf(`
						the "%s" tag's element [%s] parse params in body err,please check the request json. `,
						"qval", IN_BODY))
				}
				tfv = reflect.ValueOf(o).Elem()
			}
		case 2:
			key := tagElems[0]
			keyIn := tagElems[1]

			var kv string

			switch keyIn {
			case IN_PATH:
				vars := mux.Vars(req)
				kv = vars[key]
			case IN_QUERY:
				kv = ParseQueryGet(req, key)
			}
			switch tftk {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
				res, err := govalidator.ToInt(kv)
				if err != nil {
					panic(err)
				}
				tfv = reflect.ValueOf(res)
			case reflect.Bool:
				res, err := govalidator.ToBoolean(kv)
				if err != nil {
					panic(err)
				}
				tfv = reflect.ValueOf(res)
			case reflect.Float32, reflect.Float64:
				res, err := govalidator.ToFloat(kv)
				if err != nil {
					panic(err)
				}
				tfv = reflect.ValueOf(res)
			case reflect.String:
				tfv = reflect.ValueOf(kv)
			default:
				panic(common.ErrInternalServer(errors.New("don't support the type of filed:" + tft.Kind().String())))
			}
		}
		v.FieldByName(tf.Name).Set(tfv)
	}

	if noTagStruct {
		o := nt.Interface()
		common.Unmarshal2Object(req, &o)
		v = reflect.ValueOf(o).Elem()
	}

	return v
}
