package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	// Get请求，绑定函数
	beego.Get("/", func(context *context.Context) {
		// 获取请求参数
		name := context.Request.FormValue("name")
		//name := context.Input.Query("name")
		// 响应数据
		context.Output.Context.WriteString(fmt.Sprintf("您输入的名字是：%s", name))
	})

	// Post请求，绑定函数
	beego.Post("/", func(c *context.Context) {
		//name := c.Request.PostFormValue("name")
		name := c.Input.Query("name")
		c.Output.Context.WriteString(fmt.Sprintf("(post)您输入的名字是：%s", name))
	})

	beego.Any("/any/", func(c *context.Context) {
		c.Output.Context.WriteString(fmt.Sprintf("(%s)您输入的名字是：%s", c.Request.Method, c.Request.FormValue("name")))
	})

	beego.Run("127.0.0.1:9999")
}
