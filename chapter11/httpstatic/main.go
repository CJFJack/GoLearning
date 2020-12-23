package main

import "net/http"

func main() {
	// 读取与请求路径同名的文件夹
	//http.Handle("/www/", http.FileServer(http.Dir(".")))

	// 读取与请求路径不同名的文件夹
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./www"))))

	// 启动web服务器
	http.ListenAndServe(":8888", nil)
}
