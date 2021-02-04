package main

import (
	"github.com/astaxie/beego"
	"strings"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Data["name"] = "cjf"
	c.Data["body"] = true
	c.Data["score"] = []int{1, 2, 3, 4, 5}
	c.Data["users"] = map[int]string{1: "cjf", 2: "lulu"}
	c.Data["content"] = "abc.Abc"
	c.Data["templateInTheHtml"] = "templateInTheHtml"
	c.TplName = "index.html"
}

func (c *HomeController) Home() {
	// 若不响应任何数据，则默认找：views/控制机名称小写/方法名.tpl 渲染
}

func main() {
	beego.AddFuncMap("lower", func(in string) string {
		return strings.ToLower(in)
	})
	beego.AutoRouter(&HomeController{})

	beego.Run()
}
