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
	// 与服务端建立连接
	addr := "127.0.0.1:6666"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for i := 0; i < 4; i++ {
		// 向服务端发送数据
		fmt.Fprintf(conn, "Time: %s\n", time.Now().Format(dateFormat))
		// 接收回复数据
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		// 接收到服务器端响应
		log.Printf("接收到服务器端响应：%s", string(line))
	}
	fmt.Fprintln(conn, "quit")

}
