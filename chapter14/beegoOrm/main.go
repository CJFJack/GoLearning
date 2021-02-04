package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID        int        `orm:"column(id)"`
	Name      string     `orm:"type:varchar(20);description(姓名)"`
	Password  string     `orm:"type:varchar(1024);description(密码)"`
	Gender    bool       `orm:"default(true)"`
	Addr      string     `orm:"null;type(text)"`
	Tel       string     `orm:"size(13);index"`
	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdatedAt *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null"`
}

// 自定义表名
func (m *User) TableName() string {
	return "user"
}

// 创建索引
func (m *User) TableIndex() [][]string {
	return [][]string{
		{"Name", "Tel"},
	}
}

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=PRC&parseTime=true",
	"root", "cmdb", "127.0.0.1", "3306", "testbeegoorm",
)

func main() {
	// 注册mysql驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册dsn
	orm.RegisterDataBase("default", "mysql", dsn)
	// 注册Model
	orm.RegisterModel(new(User))

	// 执行命令
	orm.RunCommand()

	// 同步数据库
	//orm.RunSyncdb("default", true, true)

	orm.Debug = true
	// 增加
	user := &User{
		Name:     "cjf",
		Password: "123123",
		Gender:   true,
		Tel:      "13121214111",
	}

	// 数据库连接
	ormer := orm.NewOrm()

	// 插入数据
	//fmt.Printf("%#v\n", user)
	//id, err := ormer.Insert(user)
	//fmt.Println(id, err)
	//fmt.Printf("%#v\n", user)
	// 批量插入
	var users []*User
	for i := 0; i < 20; i++ {
		user = &User{
			Name:     fmt.Sprintf("cjf_%d", i),
			Password: "123123",
			Gender:   true,
			Tel:      fmt.Sprintf("1314561214%d", i),
		}
		users = append(users, user)
	}
	ormer.InsertMulti(10, users)

	// 查，Read只能查询一条数据
	user = &User{ID: 1}
	err := ormer.Read(user)
	fmt.Println(user, err)

	user = &User{Name: "cjf_0"}
	err = ormer.Read(user, "Name")
	fmt.Println(user, err)

	// 更新
	user = &User{ID: 6}
	err = ormer.Read(user)
	fmt.Println(user, err)
	user.Name = "lulu"
	// 全表字段更新
	ormer.Update(user)
	// 更新指定字段，避免同时操作同一行数据时，覆盖数据
	ormer.Update(user, "Name")

	// 删除
	user = &User{ID: 11}
	rows, err := ormer.Delete(user)
	fmt.Println(rows, err)

	// 不存在则更新
	user2 := &User{
		Name:     "lulu2",
		Password: "123123",
		Gender:   true,
		Tel:      "13121214111",
	}
	ormer.ReadOrCreate(user2, "Name")

	// 查询结果集
	querySet := ormer.QueryTable(&User{})
	fmt.Println(querySet.Count())
	var users2 []*User
	querySet.All(&users2)
	fmt.Println(users2)
	for idx, value := range users2 {
		fmt.Println(idx, value)
	}
	// filter
	// =([i]exact)  <(lt)  >(gt)  >=(gte)  <=(lte)
	// in(in)
	// like %xxx%([i]contains) start%([i]startswith)  %end([i]endswith)
	fmt.Println(querySet.Filter("name__iexact", "Cjf_19").Count())
	fmt.Println(querySet.Filter("name__icontains", "Cjf_19").Count())
	fmt.Println(querySet.Filter("name__startswith", "Cjf_19").Count())
	fmt.Println(querySet.Filter("name__endswith", "Cjf_19").Count())
	fmt.Println(querySet.Filter("id__in", []int{1, 2, 3}).Count())
	fmt.Println(querySet.Filter("id__gt", 3).Filter("id__lt", 100).Count())

	// exclude
	fmt.Println(querySet.Exclude("id__in", []int{1, 2, 3}).Count())

	// 分页
	var user3 []*User
	querySet.Limit(3).Offset(2).All(&user3)
	for idx, value := range user3 {
		fmt.Println(idx, value)
	}

	// 排序
	var user4 []*User
	querySet.OrderBy("name", "Tel").All(&user4)
	querySet.OrderBy("-name", "Tel").One(&user4)

	// And Or
	cond := orm.NewCondition()
	andCond := orm.NewCondition()
	orCond := orm.NewCondition()

	andCond = andCond.And("Name__icontains", "cjf_12")
	orCond = orCond.Or("Tel__icontains", "131").Or("Tel__icontains", "132")
	cond = cond.AndCond(andCond).AndCond(orCond)

	var users5 []*User
	querySet.SetCond(cond).All(&users5)
	for idx, value := range users5 {
		fmt.Println(idx, value)
	}

	// querySet update
	querySet.Filter("id__lt", 10).Update(orm.Params{"tel": "135711111"})
	// 原有基础上加减乘除
	querySet.Filter("id__lt", 10).Update(orm.Params{
		"tel": orm.ColValue(orm.ColAdd, 10),
	})

	// querySet delete
	querySet.Filter("id__gt", 100).Delete()

	// 原始sql
	db, err := orm.GetDB("default")
	db.Ping()
	//db.QueryRow("select * from users").Scan()
	//db.Exec()
	rawseter := ormer.Raw("insert into xxx select *")
	rawseter.Exec()

	rawseter = ormer.Raw("select * from xxx")
	var user6 *User
	rawseter.QueryRow(&user6)
	var users6 []*User
	rawseter.QueryRows(&users6)

	// 事务
	ormer.Begin()
	ormer.Commit()
	ormer.Rollback()

}
