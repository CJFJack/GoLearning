package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.String("mysql::User"),
		beego.AppConfig.String("mysql::Password"),
		beego.AppConfig.String("mysql::Host"),
		beego.AppConfig.String("mysql::Port"),
	)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("db open err: %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("db ping err: %s", err)
	}
}
