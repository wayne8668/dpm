package repositories

import (
	// "time"
	"dpm/common"
	"dpm/models"
	"encoding/json"
	"fmt"
	// "log"
)

type ProjectRepository struct{}

var (
	pr = &ProjectRepository{}
)

//单例构造函数
func NewProjectRepository() *ProjectRepository {
	return pr
}

func JSON2Map(reqJSON string) map[string]interface{} {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(reqJSON), &m); err != nil {
		//
	}
	return m
}

//新增项目
func (this *ProjectRepository) Create(reqJSON string) (numResult int64, err error) {
	fmt.Println("ProjectRepository create method invork...")
	sqlStr := "CREATE (p:PROJECT {id:{id}, name:{name}, desc:{desc},create_time:{create_time}})"
	m := JSON2Map(reqJSON)
	m["id"] = NewUUID()
	m["create_time"] = common.NowStringFormat()

	numResult, err = baseRep.execNeo(sqlStr, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("numResult", numResult)
	return numResult, err
}

//更新项目
func (this *ProjectRepository) Update(reqJSON string) (numResult int64, err error) {
	fmt.Println("ProjectRepository create method invork...")
	connStr := "MATCH (n:PROJECT) where n.id={id} set n.name ={name},n.desc = {desc}"
	m := JSON2Map(reqJSON)
	numResult, err = baseRep.execNeo(connStr, m)
	return numResult, err
}

//删除项目
func (this *ProjectRepository) Delete(reqJSON string) (numResult int64, err error) {
	fmt.Println("ProjectRepository create method invork...")
	connStr := "MATCH (n:PROJECT) where n.id={id} delete n"
	m := JSON2Map(reqJSON)
	numResult, err = baseRep.execNeo(connStr, m)
	return numResult, err
}

//添加项目分组
func (this *ProjectRepository) AddProjectTaskGroup(projectId string, tg models.TaskGroup) (numResult int64, err error) {
	connStr := `
		MATCH 
			(p:PROJECT) 
		where 
			p.pid={pid} 
		create 
			(:TASK_GROUP{tgid:{tgid},name:{name},desc:{desc}}) <- [:PROJECT_HAS_TASKGROUP] - (p)`

	m := map[string]interface{}{
		"pid":  projectId,
		"tgid": NewUUID(),
		"name": tg.Name,
		"desc": tg.Desc,
	}
	numResult, err = baseRep.execNeo(connStr, m)
	fmt.Println("ProjectRepository create method invork...")
	return numResult, err
}

//添加分组任务列表
func (this *ProjectRepository) AddProjectTaskList(tgid string, tl models.TaskList) (numResult int64, err error) {
	connStr := `MATCH (tg:TASKGROUP) 
				where tg.tgid={tgid} 
				create 
				(:TASKLIST{tlid:{tgid},name:{name},order:{order}}) <- [:TASKGROUP_HAS_TASKLIST] - (tg)`

	m := map[string]interface{}{
		"tgid":  tgid,
		"tlid":  NewUUID(),
		"name":  tl.Name,
		"order": tl.Order,
	}
	numResult, err = baseRep.execNeo(connStr, m)
	fmt.Println("ProjectRepository create method invork...")
	return numResult, err
}

//添加列表任务
func (this *ProjectRepository) AddListTask(tlid string, t models.Task) (numResult int64, err error) {
	connStr := `
	MATCH 
		(tl:TASKLIST) 
	where 
		tl.tlid={tlid} 
	create 
		(:TASK{tid:{tid},name:{name},desc:{desc},create_time:{createTime},finish_time:{finishTime},dateline:{dateline},remind_time:{remindTime}}) <- [:TASKLIST_HAS_TASK] - (tl)`

	m := map[string]interface{}{
		"tlid":       tlid,
		"tid":        NewUUID(),
		"name":       t.Name,
		"desc":       t.Desc,
		"createTime": t.CreateTime,
		"finishTime": t.FinishTime,
		"dateline":   t.DateLine,
		"remindTime": t.RemindTime}

	numResult, err = baseRep.execNeo(connStr, m)
	fmt.Println("ProjectRepository create method invork...")

	return numResult, err
}

func (this *ProjectRepository) ListA() (results []models.Project) {
	conn := GetConn()
	data, rowsMetadata, _, _ := conn.QueryNeoAll("match (p:PROJECT) return p.name,p.pid,p.desc", nil)
	fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))                            // COLUMNS: n.foo,n.bar
	fmt.Printf("FIELDS: %s %s %s\n", data[0][0].(string), data[0][1].(string), data[0][2].(string)) // FIELDS: 1 2.2
	fmt.Printf("FIELDS: %s %s %s\n", data[1][0].(string), data[1][1].(string), data[1][2].(string)) // FIELDS: 1 2.2

	results = make([]models.Project, len(data))
	for idx, row := range data {
		results[idx] = models.Project{
			Id:   row[0].(string),
			Name: row[1].(string),
			Desc: row[2].(string),
		}
	}

	//不要用map，map返回顺序可能会乱
	// r := make(map[models.Project]interface{})
	// for _, row := range data {
	// 	r[models.Project{
	// 		Pid:  row[0].(string),
	// 		Name: row[1].(string),
	// 		Desc: row[2].(string),
	// 	}] = nil
	// }

	return results
}
