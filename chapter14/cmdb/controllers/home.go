package controllers

import (
	"cmdb/base/controllers/auth"
)

type HomeController struct {
	auth.LayoutController
}

func (c *HomeController) Index() {
	c.Data["title"] = "首页"
	c.TplName = "home/index.html"
}
