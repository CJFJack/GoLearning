package main

import (
	"io"
	"net/http"
)

// 1. 定义处理器结构体
// 2. 定义处理器结构体方法ServeHTTP
// 2. 绑定url和结构体方法的关系
// 3. 启动web服务

/*
	1. 定义处理器结构体
*/
type TimeHandler struct{}

/*
	2. 定义处理器结构体方法ServeHTTP
	http.Request
	http.ResponseWriter
*/
func (t *TimeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "hello world")
}

func main() {
	/*
		3. 绑定url和处理器函数
		http.Handle(path, 处理器结构体实例指针)
	*/
	http.Handle("/time/", &TimeHandler{})

	/*
		3. 启动web服务
		http.ListenAndServe(addr, nil)
	*/
	http.ListenAndServe("127.0.0.1:8888", nil)
}
