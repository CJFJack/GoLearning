package main

import (
	"encoding/xml"
	"github.com/astaxie/beego"
)

type OutputController struct {
	beego.Controller
}

func (c *OutputController) CString() {
	c.Ctx.WriteString("Context WriteString")
}

func (c *OutputController) CBody() {
	c.Ctx.WriteString("Body WriteString")
}

func (c *OutputController) Tpl() {
	c.TplName = "tpl.html"
}

func (c *OutputController) Json() {
	c.Data["json"] = map[string]interface{}{"a": 1, "b": "2"}
	c.ServeJSON()
}

func (c *OutputController) Yaml() {
	c.Data["yaml"] = map[string]interface{}{"a": 1, "b": "2"}
	c.ServeYAML()
}

func (c *OutputController) Xml() {
	c.Data["xml"] = struct {
		XMLName xml.Name `xml:"root"`
		Name    string   `xml:"name"`
		Addr    string   `xml:"addr"`
	}{Name: "cjf", Addr: "127.0.0.1"}
	c.ServeXML()
}

func (c *OutputController) Redir() {
	c.Redirect("https://www.baidu.com", 302)
	c.Ctx.Input.Is()
}

func (c *OutputController) Ab() {
	c.Abort("404")
}

func main() {
	beego.AutoRouter(&OutputController{})

	beego.Run()
}
