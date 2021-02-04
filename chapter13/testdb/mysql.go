package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:cmdb@tcp(127.0.0.1:3308)/cmdb61?charset=utf8mb4&loc=PRC&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	fmt.Print(db, err)
	fmt.Println(db.Ping())
	result, _ := db.Exec("select * from aliyun_ecs")
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	rows, _ := db.Query("select id,instance_name from aliyun_ecs where id = ?", 46)
	var (
		id           int
		instanceName string
	)
	for rows.Next() {
		rows.Scan(&id, &instanceName)
		fmt.Println(id, instanceName)
	}
	db.Close()
}
