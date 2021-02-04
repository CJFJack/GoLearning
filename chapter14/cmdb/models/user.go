package models

import (
	"cmdb/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

// User 用户对象
type User struct {
	ID         int    `orm:"column(id)"`
	StaffID    string `orm:"column(staff_id);size(32)"`
	Name       string `orm:"size(64)"`
	NickName   string `orm:"size(64)"`
	Password   string `orm:"size(1024)"`
	Gender     int    `orm:"description:'1男0女'"`
	Tel        string `orm:"size(32)"`
	Addr       string `orm:"size(128)"`
	Email      string `orm:"size(64)"`
	Department string `orm:"size(128)"`
	Status     int
	CreatedAt  *time.Time `orm:"auto_now_add"`
	UpdatedAt  *time.Time `orm:"auto_now"`
	DeletedAt  *time.Time `orm:"null"`
	Deleted    int
}

// 注册model
func init() {
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

// 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	//return u.Password == utils.Md5Text(password)
	return utils.CheckPassword(u.Password, password)
}

func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	} else {
		return "男"
	}
}

func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	default:
		return "未知"
	}
}
