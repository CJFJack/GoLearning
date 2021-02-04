package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type InputController struct {
	beego.Controller
}

func (c *InputController) Json() {
	c.Ctx.Input.CopyBody(10 * 1024 * 1024)
	var m map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	fmt.Printf("%#v\n", m)
	fmt.Println(m["a"])
}

func main() {
	beego.AutoRouter(&InputController{})

	beego.Run()
}
