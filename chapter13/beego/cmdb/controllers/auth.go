package controllers

import (
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
	"github.com/astaxie/beego"
	"net/http"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	// get请求，加载野蛮
	// post请求，数据验证
	// 验证成功
	// 验证失败
	form := &forms.LoginForm{}
	errs := errors.New()
	if c.Ctx.Request.Method == "POST" {
		if err := c.ParseForm(form); err == nil {
			user := models.GetUserByName(form.Name)
			if user == nil {
				// 用户不存在
				errs.Add("default", "用户名或密码错误")
			} else if user.ValidPassword(form.Password) {
				// 用户密码正确
				c.Redirect("/home/index/", http.StatusFound)
			} else {
				// 用户密码不正确
				errs.Add("default", "用户名或密码错误")
			}
		} else {
			errs.Add("default", "用户名或密码错误")
		}
	}

	c.Data["form"] = form
	c.Data["errors"] = errs

	// 定义加载页面
	c.TplName = "auth/login.html"
}
