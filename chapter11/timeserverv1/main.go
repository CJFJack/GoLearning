package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 1. 定义处理器、处理器函数
	// 2. 绑定url和处理器函数的关系
	// 3. 启动web服务

	/*
		1. 处理器函数
		http.Request
		http.ResponseWriter
	*/
	TimeFunc := func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(writer, now)
	}

	/*
		2. 绑定url和处理器函数
		http.HandleFunc(path, 处理函数)
	*/
	http.HandleFunc("/time/", TimeFunc)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ctx, err := ioutil.ReadFile("index.html")
		if err == nil {
			fmt.Fprintf(writer, string(ctx))
		} else {
			fmt.Fprintf(writer, "欢迎")
		}
	})

	/*
		3. 启动web服务
		http.ListenAndServe(addr, nil)
	*/
	http.ListenAndServe("127.0.0.1:8888", nil)

}
