package controllers

import (
	"cmdb/base/controllers/base"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"html/template"
	"net/http"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	// get请求，加载页面
	// post请求，数据验证
	// 验证成功
	// 验证失败

	// 如果已经登录，访问登录页，直接跳转到指定页面
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessionUser := c.GetSession(sessionKey)
	if sessionUser != nil {
		action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
		c.Redirect(beego.URLFor(action), http.StatusFound)
	}

	form := &forms.LoginForm{}
	errs := errors.New()
	if c.Ctx.Request.Method == "POST" {
		if err := c.ParseForm(form); err == nil {
			user := services.UserService.GetByName(form.Name)
			if user == nil {
				// 用户不存在
				errs.Add("default", "用户名或密码错误")
				logs.Error(fmt.Sprintf("用户不存在： %s", form.Name))
			} else if user.ValidPassword(form.Password) {
				logs.Info(fmt.Sprintf("用户认证成功： %s", form.Name))
				// 用户密码正确
				// 设置session
				sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
				c.SetSession(sessionKey, user.ID)
				// 跳转
				action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
				c.Redirect(beego.URLFor(action), http.StatusFound)
			} else {
				// 用户密码不正确
				errs.Add("default", "用户名或密码错误")
				logs.Error(fmt.Sprintf("用户名或密码错误： %s", form.Name))
			}
		} else {
			errs.Add("default", "用户名或密码错误")
		}
	}

	c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
	c.Data["form"] = form
	c.Data["errors"] = errs

	// 定义加载页面
	c.TplName = "auth/login.html"
}

// 退出登录
func (c *AuthController) Logout() {
	c.DestroySession()
	action := beego.AppConfig.DefaultString("auth::LogoutAction", "AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusAccepted)
}
