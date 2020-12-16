package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/login", 301)
	})

	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "登录页面")
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "登录页面")
	})

	http.ListenAndServe(":8888", nil)
}
