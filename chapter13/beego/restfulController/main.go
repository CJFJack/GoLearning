package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

type HostController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Ctx.WriteString("Get")
}

func (c *HostController) Get() {
	c.Ctx.WriteString(fmt.Sprintf("Host id: %s", c.Ctx.Input.Param(":id")))
}

func main() {
	beego.Router("/home/", &HomeController{})
	beego.Router("/host/:id(\\d+)/", &HostController{})

	beego.Run()
}
