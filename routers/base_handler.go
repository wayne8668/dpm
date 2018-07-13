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

var ()

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

		fmt.Println(fv, ft)
		//返回函数输入参数的总数
		numIn := ft.NumIn()

		in := make([]reflect.Value, numIn)

		//迭代每一个参数
		for i := 0; i < numIn; i++ {
			fmt.Println("=====================================", reflect.TypeOf(w).Elem())
			//获取参数类型
			fit := ft.In(i)
			fmt.Println(fit)
			var v reflect.Value
			if fit.Kind() == reflect.Interface {
				if reflect.TypeOf(w).Implements(fit) {
					in[i] = reflect.ValueOf(w)
					continue
				}
			}

			if fit == reflect.TypeOf(req) {
				in[i] = reflect.ValueOf(req)
				continue
			}

			if fit.Kind() == reflect.Struct {
				v = structEvaluate(fit, w, req)
			}

			fmt.Println(v)
			in[i] = v
			_, err := govalidator.ValidateStruct(in[i].Interface())
			if err != nil {
				panic(common.ErrTrace(err))
			}
		}
		fmt.Println(len(in))
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
		for idx, item := range outs {
			itemInf := item.Interface()
			if yes, err := isErr(itemInf); yes {
				panic(common.ErrTrace(err))
			}
			outSli[idx] = itemInf
		}
		common.JsonResponseOK(w, &outSli)
	}
}

func structEvaluate(t reflect.Type, w http.ResponseWriter, req *http.Request) reflect.Value {
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
				"rval", IN_PATH, IN_QUERY, IN_BODY, "+", "-"))
		case 1:
			eleFst := tagElems[0]
			switch eleFst {
			case "-":
				continue
			case "+":
				if tftk != reflect.Struct {
					panic(common.ErrInternalServerf(`
						the "%s" tag's element [%s] support struct filed only. `,
						"rval", "+"))
				}
				tfv = structEvaluate(tft, w, req)
			case IN_BODY:
				if tftk != reflect.Struct {
					panic(common.ErrInternalServerf(`
						the "%s" tag's element [%s] support struct filed only. `,
						"rval", IN_BODY))
				}
				o := reflect.New(tft).Interface()
				common.Unmarshal2Object(w, req, &o)
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
		common.Unmarshal2Object(w, req, &o)
		v = reflect.ValueOf(o).Elem()
	}

	return v
}
