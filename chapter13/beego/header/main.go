package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

const (
	cookieKey = "abc123"
)

type HeaderController struct {
	beego.Controller
}

func (c *HeaderController) Cookie() {
	//name := c.Ctx.Input.Cookie("name")
	name := c.Ctx.GetCookie("name")
	fmt.Println(name)

	c.Ctx.SetCookie("vi", "1235")
}

func (c *HeaderController) SecureCookie() {
	test, b := c.Ctx.GetSecureCookie(cookieKey, "test")
	fmt.Println(test, b)
	c.Ctx.SetSecureCookie(cookieKey, "test", "xxxxx")
}

func main() {
	beego.AutoRouter(&HeaderController{})

	beego.Run()
}
