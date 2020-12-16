package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {

			// 方式一
			// 解析参数
			request.ParseForm()
			// Get获取参数，当参数不存在时，返回空字符串，当参数有多个字时，只能获取一个值
			fmt.Println(request.Form.Get("a"))
			// 当参数的值有多个时，使用map方法获取
			fmt.Println(request.Form["a"])

			// 方式二
			//fmt.Println(request.FormValue("a"))
		}
	})

	http.ListenAndServe("127.0.0.1:7777", nil)

}
