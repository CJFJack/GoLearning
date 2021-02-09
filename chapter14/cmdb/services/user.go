package services

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

type userService struct {
}

// 根据user id 查询用户信息
func (s *userService) GetByPk(pk int) *models.User {
	user := &models.User{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

// 新增用户
func (s *userService) Add(form *forms.UserAddForm) {
	ormer := orm.NewOrm()
	if user := s.GetByName(form.Name); user != nil {
		user.StaffID = form.StaffID
		user.Deleted = 0
		user.Status = 0
		user.NickName = form.NickName
		user.Password = utils.GeneratePassword(form.Password)
		user.Gender = form.Gender
		user.Tel = form.Tel
		user.Email = form.Email
		user.Department = form.Department
		ormer.Update(user, "Deleted", "Status", "StaffID", "NickName", "Password", "Gender", "Tel", "Email", "Department")
		return
	}
	user := &models.User{
		StaffID:    form.StaffID,
		Name:       form.Name,
		NickName:   form.NickName,
		Password:   utils.GeneratePassword(form.Password),
		Gender:     form.Gender,
		Tel:        form.Tel,
		Email:      form.Email,
		Department: form.Department,
		Status:     0,
	}
	ormer.ReadOrCreate(user, "Name")
}

// 修改用户信息
func (s *userService) Modify(form *forms.UserModifyForm) {
	if user := s.GetByPk(form.ID); user != nil {
		user.Name = form.Name
		ormer := orm.NewOrm()
		ormer.Update(user, "Name")
	}
}

// 删除用户数据
func (s *userService) Delete(pk int) {
	if user := s.GetByPk(pk); user != nil {
		now := time.Now()
		user.DeletedAt = &now
		user.Deleted = 1
		ormer := orm.NewOrm()
		ormer.Update(user, "DeletedAt", "Deleted")
	}
}

// 修改用户密码
func (s *userService) ModifyPassword(pk int, password string) {
	if user := s.GetByPk(pk); user != nil {
		user.Password = utils.GeneratePassword(password)
		orm := orm.NewOrm()
		orm.Update(user, "Password")
	}

}

// 通过用户名获取用户指针
func (s *userService) GetByName(name string) *models.User {
	ormer := orm.NewOrm()
	user := &models.User{Name: name}
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// 查询用户
func (s *userService) Query(q string) []*models.User {
	var users []*models.User
	querySet := orm.NewOrm().QueryTable(&models.User{})
	cond := orm.NewCondition()
	if q != "" {
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("tel__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("email__icontains", q)
		cond = cond.Or("department__icontains", q)
	}
	cond = cond.AndNot("deleted__exact", 1)
	querySet = querySet.SetCond(cond)
	querySet.All(&users)
	return users
}

// 用户操作实例
var UserService = new(userService)
