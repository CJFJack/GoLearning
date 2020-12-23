package main

import (
	"fmt"
	"regexp"
)

func main() {

	// 匹配字符串
	match, err := regexp.MatchString("^(135|158)\\d{8}$", "13512215112")
	if err == nil {
		fmt.Println(match)
	} else {
		fmt.Println(err)
	}

	match, err = regexp.MatchString("^[a-zA-Z0-9]{1,32}@[a-z]{1,12}[.]edu$", "abc@admin.edu")
	if err == nil {
		fmt.Println(match)
	} else {
		fmt.Println(err)
	}

	// 生成正则结构体指针
	// 匹配
	reg, err := regexp.Compile("^(135|158)\\d{8}$")
	fmt.Println(reg.MatchString("13512215112"))
	// 替换
	reg, err = regexp.Compile("(135|158)\\d{8}")
	fmt.Println(reg.ReplaceAllString("我的手机号码是13512215112", "135********"))
	// 查找
	fmt.Println(reg.FindAllString("我的手机号码是13512215112,13141211214,13513121121", -1))
	// 分割
	reg, err = regexp.Compile("[,;:]")
	fmt.Println(reg.Split("我的手机号码是13512215112,13141211214,13513121121;13sdfasdfs:ada2sdafasdf", 3))

}
