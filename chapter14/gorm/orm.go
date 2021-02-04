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
	UserId   int    `gorm:"primary_key"`
	Name     string `gorm:"size:10;default:'cjf';comment:'姓名'"`
	Password string `gorm:"not null"`
	Birthday time.Time
	Sex      bool
	Tel      string
	Addr     string
	Desc     string
	Score    float32 `gorm:"unique"`
}

func (t *User2) TableName() string {
	return "user_second"
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true", "root", "cmdb", "127.0.0.1", 3308, "testgorm")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("gorm open err: %s", err)
	}
	defer db.Close()

	//fmt.Println(db.HasTable("users"))
	//fmt.Println(db.HasTable(&User{}))
	//
	//fmt.Println(db.CreateTable(&User{}, &User2{}))

	db.DropTable(&User2{})
	fmt.Println(db.AutoMigrate(&User2{}))
	//db.Model(&User{}).AddIndex("idx_name","name(10)")
	//db.Model(&User{}).RemoveIndex("a")

	//db.Model(&User{}).ModifyColumn("birthday", "date")
	//db.Model(&User{}).DropColumn("score")

	db.Close()
}
