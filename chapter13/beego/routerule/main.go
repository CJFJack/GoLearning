package main

import (
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Query() {
	c.Ctx.WriteString("add")
}

func main() {
	beego.Router("/task/", &TaskController{}, "get:Query")

	beego.Run()
}
