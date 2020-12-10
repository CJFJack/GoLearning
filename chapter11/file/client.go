package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func ls(conn net.Conn) {
	fmt.Fprintf(conn, "ls|0|")
	reader := bufio.NewReader(conn)
	sizeText, _ := reader.ReadString('|')
	size, _ := strconv.Atoi(sizeText[:len(sizeText)-1])
	for size > 0 {
		name, _ := reader.ReadString(':')
		fmt.Println(name[:len(name)-1])
		size--
	}
}

func cat(conn net.Conn, filename string) {
	fmt.Fprintf(conn, "cat|1|%s|", filename)

	reader := bufio.NewReader(conn)
	sizeText, _ := reader.ReadString('|')
	size, _ := strconv.Atoi(sizeText[:len(sizeText)-1])
	if size > 0 {
		ctx := make([]byte, size)
		n, _ := reader.Read(ctx)
		fmt.Println(string(ctx[:n]))
	} else {
		fmt.Println("文件为空")
	}

}

func main() {
	// 定义日志文件
	logfile, _ := os.OpenFile("client.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	log.SetOutput(logfile)
	logfile.Close()

	addr := "127.0.0.1:6666"
	// 与服务端建立连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("connect err: %s\n", err)
	}
	defer conn.Close()
	log.Println("connected fileserver..")

	scanner := bufio.NewScanner(os.Stdin)

END:
	for {
		fmt.Print("请输入指令：")
		scanner.Scan()
		input := scanner.Text()
		cmds := strings.Split(input, " ")
		switch cmds[0] {
		case "ls":
			ls(conn)
		case "cat":
			cat(conn, cmds[1])
		case "put":
		case "download":
		case "delete":
		case "quit":
			break END
		default:
			fmt.Println("指令错误，请重新输入")
		}
	}

}
