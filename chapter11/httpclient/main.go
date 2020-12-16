package main

import (
	"fmt"
	"net/http"
)

func main() {

	// get请求
	response, err := http.Get("http://localhost:8888")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Proto, response.Status)
		ctx := make([]byte, 1024)
		n, _ := response.Body.Read(ctx)
		fmt.Println(string(ctx[:n]))
		fmt.Println(response.Header)
	}

	// post请求 application/json
	//buffer := bytes.NewBufferString(`{"a": 1, "b":2}`)
	//response, err = http.Post("http://localhost:8888", "application/json", buffer)
	//fmt.Println(response.Proto, response.Status)
	//fmt.Println(response.Body)
	//fmt.Println(response.Header)

	// post请求 application/x-www-form-urlencoded
	//params := url.Values{}
	//params.Add("a", "1")
	//params.Add("b", "2")
	//params.Add("c", "3")
	//response, _ := http.PostForm("http://localhost:8888", params)
	//fmt.Println(response.Proto, response.Status)
	//fmt.Println(response.Body)
	//fmt.Println(response.Header)

}
