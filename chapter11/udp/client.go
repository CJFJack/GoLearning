package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Printf("connect err: %s", err)
	}
	defer conn.Close()
	log.Printf("conneted %s", addr)

	reader := bufio.NewReader(conn)
	for {
		// 发送消息
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("请输入数据：")
		scanner.Scan()
		fmt.Fprintf(conn, "%s\n", scanner.Text())
		// 接收响应
		ctx := make([]byte, 1024)
		n, err := reader.Read(ctx)
		if err == nil {
			fmt.Printf("service: %s\n", string(ctx[:n]))
		}
	}
}
