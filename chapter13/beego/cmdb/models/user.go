package models

import (
	"cmdb/utils"
)

// User 用户对象
type User struct {
	ID         int
	StaffID    string
	Name       string
	NickName   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	Status     string
}

// 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return u.Password == utils.Md5Text(password)
}

// 通过用户名获取用户指针
func GetUserByName(name string) *User {
	var user = &User{}
	if err := db.QueryRow("select id, name, password from user where name = ?", name).Scan(&user.ID, &user.Name, &user.Password); err == nil {
		return user
	}
	return nil
}
