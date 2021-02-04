package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		method := request.Method
		url := request.URL
		path := url.Path
		rawQuery := url.RawQuery

		proxyUrl := "http://127.0.0.1:7777" + path + rawQuery
		if method == "GET" {
			response, err := http.Get(proxyUrl)
			if err != nil {
				fmt.Println(err)
			}
			ctx := make([]byte, 1024)
			for {
				n, err := response.Body.Read(ctx)
				fmt.Fprintf(writer, string(ctx[:n]))
				if err == io.EOF {
					break
				}
			}
		}
		if method == "POST" {
			buffer := bytes.NewBufferString("")
			ctx := make([]byte, 10)
			for {
				n, err := request.Body.Read(ctx)
				buffer.Write(ctx[:n])
				if err != nil {
					if err == io.EOF {
						break
					}
				}
			}
			http.Post(proxyUrl, "application/inputJson", buffer)
		}
	})

	http.ListenAndServe(":8888", nil)
}
