package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		method := request.Method
		fmt.Println(request.Header)
		fmt.Println(request.RemoteAddr)
		fmt.Println(request.Body)
		fmt.Println(request.Proto)
		fmt.Println(request.URL)
		if method == "GET" {
			fmt.Fprintf(writer, "welcome")
		}

		if method == "POST" {
			var info map[string]interface{}
			decoder := json.NewDecoder(request.Body)
			decoder.Decode(&info)
			fmt.Println(info)
		}
	})

	http.ListenAndServe(":7777", nil)

}
