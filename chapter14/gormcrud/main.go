package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Birthday time.Time
	Sex      bool
	Tel      string
	Addr     string
	Desc     string
	Score    float32
}

type User2 struct {
	Name     string
	Password string
	Birthday time.Time
	Sex      bool
	Tel      string
	Addr     string
	Desc     string
	Score    float32
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true", "root", "cmdb", "127.0.0.1", 3308, "testgorm")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("gorm open err: %s", err)
	}
	db.LogMode(true)
	defer db.Close()

	//db.DropTable(&User{})
	db.AutoMigrate(&User{})

	// 新增
	//user := User{
	//	Name:     "cjf",
	//	Password: "123123",
	//	Birthday: time.Date(2020, 9, 7, 0, 0, 0, 0, time.UTC),
	//	Sex:      false,
	//	Tel:      "13414151114",
	//	Addr:     "bbc",
	//}
	//fmt.Println(db.NewRecord(user))
	//db.Create(&user)
	//fmt.Println(db.NewRecord(user))
	//if db.NewRecord(user) {
	//	db.Create(&user)
	//}

	// 查询
	//var user2 User
	//db.First(&user2, "name=?", "cjf")
	//fmt.Println(user2)
	//
	//var user3 User
	//db.Last(&user3, "name=?", "lulu")
	//fmt.Println(user3)
	//
	//var users []User
	////db.Find(&users,"name=?", "cjf")
	//db.Order("id desc").Select("name,password").Where("tel in (?)", []string{"13414151114", "13570115111"}).Find(&users)
	//fmt.Println(users)
	//
	//var count int
	//db.Model(&User{}).Count(&count)
	//fmt.Println(count)
	//
	//rows, _ := db.Model(&User{}).Select("name, password").Rows()
	//for rows.Next() {
	//	var name, password string
	//	rows.Scan(&name, &password)
	//	fmt.Println(name, password)
	//}
	//
	//rows, _ = db.Model(&User{}).Select("name,count(*)").Group("name").Having("count(*)>?", 1).Rows()
	//for rows.Next() {
	//	var name string
	//	var count int
	//	rows.Scan(&name, &count)
	//	fmt.Println(name, count)
	//}

	var user4 User
	if db.First(&user4, "name = ?", "cjf").Error == nil {
		//fmt.Println(user4)
		//user4.Addr = "abcd"
		//db.Save(&user4)
	} else {
		fmt.Println("记录不存在")
	}

	var user5 User
	if db.First(&user5, "name = ?", "cjf").Error == nil {
		fmt.Println(user5)
		db.Delete(&user5)
	} else {
		fmt.Println("记录不存在")
	}

	//db.Model(&User{}).Where("id > ?", 1).UpdateColumn("tel", "137894151")
	//db.Model(&User{}).Where("id > ?", 1).UpdateColumns(map[string]interface{}{"tel": "132456789", "addr": "china"})
	//db.Model(&User{}).Where("id > ?", 1).Updates(map[string]interface{}{"tel": "132456789", "addr": "usa"})
	//db.Model(&User{}).Where("id > ?", 1).Updates(User{Tel: "123123", Addr: "gz"})

	db.Where("id > ?", 2).Unscoped().Delete(&User{})

	db.AutoMigrate(&User2{})

	//user2 := User2{
	//	Name:     "cjf",
	//	Password: "123123",
	//	Birthday: time.Date(2020, 9, 7, 0, 0, 0, 0, time.UTC),
	//	Sex:      false,
	//	Tel:      "13414151114",
	//	Addr:     "bbc",
	//}
	//fmt.Println(db.NewRecord(user2))
	//db.Create(&user2)
	db.Delete(&User2{})

	// 执行原始sql
	insert := db.Exec("insert into users(`name`) values('feiji')")
	fmt.Println(insert.RowsAffected)

	var name string
	db.Raw("select name from users").Row().Scan(&name)
	fmt.Println(name)

	var user6 []User
	db.Raw("select * from users").Scan(&user6)
	fmt.Println(user6)

	rows, _ := db.Raw("select name,tel from users").Rows()
	for rows.Next() {
		var name, tel string
		rows.Scan(&name, &tel)
		fmt.Println(name, tel)
	}

	delete := db.Exec("delete from users")
	fmt.Println(delete.RowsAffected)

}
