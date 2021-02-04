package controllers

import "cmdb/base/controllers/base"

// 错误处理控制器
type ErrorController struct {
	base.BaseController
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
}

func (c *ErrorController) ErrorNotPermission() {
	c.TplName = "error/not_permission.html"
}
