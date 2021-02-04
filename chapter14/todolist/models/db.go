package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const (
	dbUser     string = "root"
	dbPassword string = "cmdb"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3308
	dbName     string = "todolist"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

func init() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("gorm open connect err: %s", err)
	}
	if db.DB().Ping() != nil {
		log.Panicf("db ping err")
	}
}
