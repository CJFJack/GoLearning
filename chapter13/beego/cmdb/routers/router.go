package routers

import (
	"cmdb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
}
