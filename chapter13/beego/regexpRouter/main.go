package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.Get("/id/?:id(\\d+)/", func(c *context.Context) {
		c.WriteString(fmt.Sprintf("匹配，id：%s", c.Input.Param(":id")))
	})

	beego.Get("/name/:name(\\w+)/", func(c *context.Context) {
		c.WriteString(fmt.Sprintf("匹配，name：%s", c.Input.Param(":name")))
	})

	beego.Get("/any/:any/", func(c *context.Context) {
		c.WriteString(fmt.Sprintf("匹配，any：%s", c.Input.Param(":any")))
	})

	beego.Get("/file/*", func(c *context.Context) {
		c.WriteString(fmt.Sprintf("匹配，file：%s", c.Input.Param(":splat")))
	})

	beego.Get("/ext/*.*", func(c *context.Context) {
		c.WriteString(fmt.Sprintf("匹配，ext：%s, %s", c.Input.Param(":path"), c.Input.Param(":ext")))
	})

	beego.Run()
}
