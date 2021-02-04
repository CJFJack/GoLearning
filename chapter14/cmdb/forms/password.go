package forms

import (
	"cmdb/models"
	"github.com/astaxie/beego/validation"
	"regexp"
)

// 修改密码表单
type PasswordModifyForm struct {
	LoginUser   *models.User `form:"-"`
	OldPassword string       `form:"old_password"`
	Password    string       `form:"password"`
	Password2   string       `form:"password2"`
}

// 数据检查
func (f *PasswordModifyForm) Valid(valid *validation.Validation) {
	// 验证旧密码
	if ok := f.LoginUser.ValidPassword(f.OldPassword); !ok {
		valid.AddError("default.default", "旧密码错误")
		return
	}
	// 验证密码范围，大小写英文字母、特殊字符
	passwordRegexp := "^[0-9a-zA-Z_.\\$\\!#!]{6,20}$"
	valid.Match(f.Password, regexp.MustCompile(passwordRegexp), "default.default.default").Message("密码格式不正确")
	if valid.HasErrors() {
		return
	} else if f.Password != f.Password2 {
		valid.AddError("default.default", "两次密码不一致")
	} else if f.OldPassword == f.Password {
		valid.AddError("default.default", "新旧密码不能一致")
	}
}
