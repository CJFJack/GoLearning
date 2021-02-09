package main

import (
	_ "cmdb/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/cmdb.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLevel(logs.LevelDebug)
	logs.Async()

	beego.Run()
}
