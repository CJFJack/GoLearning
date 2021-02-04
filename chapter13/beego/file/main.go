package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) File() {
	fmt.Println(c.GetFile("images"))
	c.SaveToFile("images", "upload/xxx.txt")
}

func main() {
	beego.AutoRouter(&FileController{})

	beego.Run()
}
