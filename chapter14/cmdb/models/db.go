package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=PRC&parseTime=true",
	beego.AppConfig.String("mysql::User"),
	beego.AppConfig.String("mysql::Password"),
	beego.AppConfig.String("mysql::Host"),
	beego.AppConfig.String("mysql::Port"),
	beego.AppConfig.String("mysql::Database"),
)

func init() {
	orm.Debug = true
	// 注册mysql驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册dsn
	orm.RegisterDataBase("default", "mysql", dsn)

	// 测试数据库连接
	if db, err := orm.GetDB("default"); err != nil {
		log.Fatalf("orm get db err:%s", err)
	} else if err := db.Ping(); err != nil {
		log.Fatalf("orm db ping err:%s", err)
	}
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 100)

}
