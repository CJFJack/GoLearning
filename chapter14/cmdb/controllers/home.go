package controllers

import (
	"cmdb/base/controllers/auth"
)

type HomeController struct {
	auth.LayoutController
}

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
