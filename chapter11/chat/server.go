package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := "127.0.0.1:6666"

	// 监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("listen on [%s]", addr)

	// 循环接收客户端连接
	for {
		// 接收客户端连接
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		reader := bufio.NewReader(conn)
		scanner := bufio.NewScanner(os.Stdin)
		func() {
			defer conn.Close()
			log.Printf("client[%s] is connected...", conn.RemoteAddr())
			// 读取客户端请求数据
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					log.Println(err)
					break
				} else {
					if string(line) == "quit" {
						break
					}
					log.Printf("接收到数据：%s\n", string(line))
					// 向客户端回复数据
					fmt.Print("请输入消息：")
					scanner.Scan()
					fmt.Fprintf(conn, "%s\n", scanner.Text())
				}
			}
		}()
	}
}
