package forms

// 登录表单
type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}
