package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"rpcclient/data"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 定义请求对象
	request := &data.CalculatorRequest{2, 5}

	// 定义响应对象
	response := &data.CalculatorResponse{}

	// 调用远程方法
	// 同步调用
	//err = conn.Call("Calculator.Add", request, response)
	err = conn.Call("calc.Add", request, response)
	// 获取结果
	fmt.Println(err, response.Result)

	// 异步调用
	call := conn.Go("calc.Add", request, response, nil)
	// 获取管道结果，如果管道还没有结果，则一直阻塞
	<-call.Done
	// 获取结果
	fmt.Println(response.Result)

}
