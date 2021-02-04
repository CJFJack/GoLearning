package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

type UserController struct {
	auth.LayoutController
}

// 查询用户信息
func (c *UserController) Query() {
	beego.ReadFromRequest(&c.Controller)
	q := c.GetString("q")
	c.Data["users"] = services.UserService.Query(q)
	c.Data["q"] = q
	c.TplName = "user/query.html"
}

// 新增用户
func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		form := &forms.UserAddForm{}
		if err := c.ParseForm(form); err == nil {
			services.UserService.Add(form)
			c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}
	}
	c.TplName = "user/add.html"
}

// 修改用户信息
func (c *UserController) Modify() {
	//c.Abort("NotPermission")
	//return

	form := &forms.UserModifyForm{}
	// GET 获取参数
	// POST 修改用户
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.UserService.Modify(form)
			flash := beego.NewFlash()
			flash.Set("notice", "修改用户信息成功")
			flash.Store(&c.AuthorizationController.BaseController.Controller)
			c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if user := services.UserService.GetByPk(pk); user != nil {
			form.ID = user.ID
			form.Name = user.Name
		}
	}
	c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["form"] = form
	c.TplName = "user/modify.html"
}

// 删除用户
func (c *UserController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil && c.LoginUser.ID != pk {
		services.UserService.Delete(pk)
	}
	c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}
