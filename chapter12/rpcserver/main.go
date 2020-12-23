package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpcserver/service"
	"sync"
)

func main() {
	addr := ":9999"

	// 注册服务，未指定服务名称，默认为结构体名称
	//rpc.Register(&service.Calculator{})
	// 注册服务，指定服务名称
	rpc.RegisterName("calc", &service.Calculator{})

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("[+] listen on: %s", addr)

	for {
		wg := sync.WaitGroup{}
		wg.Add(1)
		// 使用例程
		go func() {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("[-]error client: %s\n", err.Error())
				return
			}
			defer conn.Close()

			log.Printf("[+] client connected: %s\n", conn.RemoteAddr())
			// 启动jsonrpc处理客户端请求
			jsonrpc.ServeConn(conn)
			wg.Done()
		}()
		wg.Wait()
	}

}
