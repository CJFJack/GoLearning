package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	// 方式一
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	// 解析form-data
	//	request.ParseMultipartForm(1024*1024)
	//	fmt.Println(request.MultipartForm)
	//	// 获取上传的文件
	//	file, _ := request.MultipartForm.File["name"][0].Open()
	//	writeFile, _ := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	//	fmt.Println(io.Copy(writeFile, file))
	//})

	// 方式二
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 获取上传的文件
		file, header, _ := request.FormFile("name")
		fmt.Println(header.Header)
		fmt.Println(header.Size)
		fmt.Println(header.Filename)
		writeFile, _ := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		fmt.Println(io.Copy(writeFile, file))
	})

	// 启动web服务
	http.ListenAndServe(":7777", nil)

}
