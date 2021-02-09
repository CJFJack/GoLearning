package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User 用户对象
type Task struct {
	ID        int        `orm:"column(id)"`
	TaskName  string     `orm:"size(50);column(task_name)"`
	UserId    int        `orm:"column(user_id)"`
	Status    int        `orm:"column(status)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Deleted   int
}

// 注册model
func init() {
	orm.RegisterModel(new(Task))
	//orm.RunSyncdb("default", false, true)
}

func (t *Task) StatusText() string {
	switch t.Status {
	case 0:
		return "新创建"
	case 1:
		return "进行中"
	case 2:
		return "已完成"
	default:
		return "未知"
	}
}

func (t *Task) UserName() string {
	user := &User{ID: t.UserId}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user.Name
	}
	return ""
}
