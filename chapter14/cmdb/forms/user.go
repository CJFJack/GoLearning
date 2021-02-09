package forms

// 用户修改表单
type UserModifyForm struct {
	ID   int    `form:"id"`
	Name string `form:"name"`
}

// 用户修改表单
type UserAddForm struct {
	StaffID    string `form:"staff_id"`
	Name       string `form:"name"`
	NickName   string `form:"nick_name"`
	Password   string `form:"password"`
	Gender     int    `form:"gender"`
	Tel        string `form:"tel"`
	Addr       string `form:"addr"`
	Email      string `form:"email"`
	Department string `form:"department"`
	Status     int    `form:"status"`
}
