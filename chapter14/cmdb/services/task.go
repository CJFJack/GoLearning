package services

import (
	"cmdb/forms"
	"cmdb/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type taskService struct {
}

// 查询TODOLIST
func (s *taskService) Query(q string) []*models.Task {
	var tasks []*models.Task
	querySet := orm.NewOrm().QueryTable(&models.Task{})
	cond := orm.NewCondition()
	if q != "" {
		cond = cond.Or("task_name__icontains", q)
	}
	cond = cond.AndNot("deleted__exact", 1)
	querySet = querySet.SetCond(cond)
	querySet.All(&tasks)
	return tasks
}

// 新增TODOLIST
func (s *taskService) Add(form *forms.TaskAddForm) {
	ormer := orm.NewOrm()
	user := &models.Task{
		TaskName: form.TaskName,
		UserId:   form.UserId,
		Status:   form.Status,
	}
	ormer.ReadOrCreate(user, "TaskName")
}

// 根据task id 查询TODOLIST信息
func (s *taskService) GetByPk(pk int) *models.Task {
	task := &models.Task{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(task); err == nil {
		return task
	}
	return nil
}

// 修改TODOLIST信息
func (s *taskService) Modify(form *forms.TaskModifyForm) {
	if task := s.GetByPk(form.ID); task != nil {
		task.Status = form.Status
		ormer := orm.NewOrm()
		ormer.Update(task, "Status")
	}
}

// 删除TODOLIST数据
func (s *taskService) Delete(pk int) {
	if task := s.GetByPk(pk); task != nil {
		now := time.Now()
		task.DeletedAt = &now
		task.Deleted = 1
		ormer := orm.NewOrm()
		ormer.Update(task, "DeletedAt", "Deleted")
	}
}

// TODOLIST操作实例
var TaskService = new(taskService)
