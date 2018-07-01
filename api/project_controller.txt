/*
* Project 控制器
 */
package api

import (
	"dpm/models"
	"dpm/repositories"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	projectRepository = repositories.NewProjectRepository()
)

/*
* 返回我的项目
 */
func IndexProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("{router_name : project_index}"); err != nil {
		panic(err)
	}
}

/*
* 新增项目
 */
func CreateProject(w http.ResponseWriter, r *http.Request) {

	reqJSON := unmarshal2JSON(w, r)
	m := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&m)
	if numResult, err := projectRepository.Create(reqJSON); err != nil {
		m["err"] = err
	} else {
		m["numResult"] = numResult
	}
	jsonResponse(w, http.StatusCreated, m)
}

/*
* 修改项目
 */
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello update")
	params := mux.Vars(r)

	fmt.Println(params)
	id := params["id"]
	reqJSON := unmarshal2JSON(w, r)
	fmt.Println(id, reqJSON)
	json.NewEncoder(w).Encode(reqJSON)
	// m := make(map[string]interface{})
	// if numResult, err := projectRepository.Update(reqJSON); err != nil {
	// 	m["err"] = err
	// } else {
	// 	m["numResult"] = numResult
	// }
	// controllers.WriteResult(w, http.StatusCreated, m)
}

/*
* 删除项目
 */
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	reqJSON := unmarshal2JSON(w, r)
	m := make(map[string]interface{})
	if numResult, err := projectRepository.Delete(reqJSON); err != nil {
		m["err"] = err
	} else {
		m["numResult"] = numResult
	}
	jsonResponse(w, http.StatusCreated, m)
}

/*
* 新增任务分组
 */
func AddTaskGroup(w http.ResponseWriter, r *http.Request) {
	var model models.TaskGroup
	unmarshal2Object(w, r, model)
	m := make(map[string]interface{})
	if numResult, err := projectRepository.AddProjectTaskGroup("", model); err != nil {
		m["err"] = err
	} else {
		m["numResult"] = numResult
	}
	jsonResponse(w, http.StatusCreated, m)
}

/*
* 新增任务列表
 */
// func AddTaskList(w http.ResponseWriter, r *http.Request) {
// 	order, _ := this.GetInt("order")
// 	tl := models.TaskList{
// 		Name:  this.GetString("name"),
// 		Order: order,
// 	}
// 	this.AddProjectTaskList(this.GetString("tgid"), tl)
// 	this.TplName = "index.tpl"
// }
