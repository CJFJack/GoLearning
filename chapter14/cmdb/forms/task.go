package forms

// TODOLIST修改表单
type TaskModifyForm struct {
	ID       int    `form:"id"`
	TaskName string `form:"task_name"`
	Status   int    `form:"status"`
}

// TODOLIST修改表单
type TaskAddForm struct {
	StaffID  string `form:"staff_id"`
	TaskName string `form:"task_name"`
	UserId   int    `form:"user_id"`
	Status   int    `form:"status"`
}
