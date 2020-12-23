package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"

func main() {
	addr := "127.0.0.1:6666"

	// 监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("listen on [%s]", addr)

	// 循环接收
	for {
		// 接收客户端连接
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		func() {
			defer conn.Close()
			log.Printf("client[%s] is connected...", conn.RemoteAddr())
			// 读取客户端请求数据
			reader := bufio.NewReader(conn)
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
					// 回复数据
					ackMsg := fmt.Sprintf("Over：%s\n", time.Now().Format(dateFormat))
					conn.Write([]byte(ackMsg))
					fmt.Fprintf(conn, "Over：%s\n", time.Now().Format(dateFormat))
				}
			}
		}()
	}
}
