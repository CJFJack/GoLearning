package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Query() {
	fmt.Println(c.Ctx.Request.Method)
	c.Ctx.WriteString("query")
}

func main() {
	beego.AutoRouter(&TaskController{})

	beego.Run()
}
