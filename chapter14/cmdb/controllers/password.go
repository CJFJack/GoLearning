package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/services"
	"github.com/astaxie/beego/validation"
	"html/template"
)

type PasswordController struct {
	auth.LayoutController
}

// 修改用户密码
func (c *PasswordController) Modify() {
	form := &forms.PasswordModifyForm{LoginUser: c.LoginUser}
	errs := errors.New()
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证表单数据
			valid := &validation.Validation{}
			if success, err := valid.Valid(form); err != nil {
				errs.Add("default", err.Error())
			} else if !success {
				errs.AddValidation(valid)
			} else {
				services.UserService.ModifyPassword(c.LoginUser.ID, form.Password)
				c.Data["text"] = "修改成功"
				//c.DestroySession()
				//action := beego.AppConfig.DefaultString("auth::LoginAction", "AuthController.Login")
				//c.Redirect(beego.URLFor(action), http.StatusFound)
			}
		}
	}
	c.Data["xsrf_token"] = template.HTML(c.XSRFFormHTML())
	c.Data["errors"] = errs
	c.TplName = "password/modify.html"
}
