package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
	"github.com/astaxie/beego"
	"net/http"
)

type TaskController struct {
	auth.LayoutController
}

// 查询TODOLIST信息
func (c *TaskController) Query() {
	beego.ReadFromRequest(&c.Controller)
	q := c.GetString("q")
	c.Data["tasks"] = services.TaskService.Query(q)
	c.Data["q"] = q
	c.Data["title"] = "TODOLIST查询"
	c.TplName = "task/query.html"
}

// 新增TODOLIST
func (c *TaskController) Add() {
	if c.Ctx.Input.IsPost() {
		form := &forms.TaskAddForm{}
		if err := c.ParseForm(form); err == nil {
			form.UserId = c.LoginUser.ID
			services.TaskService.Add(form)
			c.Redirect(beego.URLFor("TaskController.Query"), http.StatusFound)
		}
	}
	c.TplName = "task/add.html"
}

// 修改TODOLIST信息
func (c *TaskController) Modify() {
	//c.Abort("NotPermission")
	//return

	form := &forms.TaskModifyForm{}
	// GET 获取参数
	// POST 修改TODOLIST
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.TaskService.Modify(form)
			flash := beego.NewFlash()
			flash.Set("notice", "修改TODOLIST信息成功")
			flash.Store(&c.AuthorizationController.BaseController.Controller)
			c.Redirect(beego.URLFor("TaskController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if task := services.TaskService.GetByPk(pk); task != nil {
			form.ID = task.ID
			form.TaskName = task.TaskName
			form.Status = task.Status
		}
	}
	c.Data["form"] = form
	c.Data["title"] = "TODOLIST编辑"
	c.TplName = "task/modify.html"
}

// 删除TODOLIST
func (c *TaskController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.TaskService.Delete(pk)
	}
	c.Redirect(beego.URLFor("TaskController.Query"), http.StatusFound)
}
