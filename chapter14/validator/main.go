package main

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func main() {
	valid := &validation.Validation{}
	text := ""

	// 验证
	valid.Required(text, "required.required.required")
	valid.Alpha("123abc", "alpha.alpha.alpha")
	valid.Tel("135xxxxxxxx", "tel.tel.tel")
	valid.Max("135xxxxxxxx", 3, "max.max.max")

	// 获取验证结果
	if valid.HasErrors() {
		fmt.Println(valid.Errors)
		fmt.Println(valid.ErrorsMap)
	}

}
