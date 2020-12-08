package main

import (
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

	// 向服务端发送数据
	fmt.Fprintf(conn, "Time1: %s\n", time.Now().Format(dateFormat))
	fmt.Fprintf(conn, "Time2: %s\n", time.Now().Format(dateFormat))
	fmt.Fprintf(conn, "Time3: %s\n", time.Now().Format(dateFormat))

}
