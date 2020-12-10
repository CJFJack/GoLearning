package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func cmd(conn net.Conn) (string, []string) {
	reader := bufio.NewReader(conn)
	op, err := reader.ReadString('|')
	if err != nil {
		if err == io.EOF {
			return "quit", []string{}
		}
		fmt.Printf("read cmd err: %s", err)
	}
	cntText, _ := reader.ReadString('|')
	cnt, _ := strconv.Atoi(cntText[:len(cntText)-1])
	args := make([]string, 0, cnt)
	for cnt > 0 {
		param, _ := reader.ReadString('|')
		args = append(args, param[:len(param)-1])
		cnt--
	}
	return op[:len(op)-1], args
}

func ls(conn net.Conn) {
	file, _ := os.Open(".")
	names, _ := file.Readdirnames(-1)
	suffix := ""
	if len(names) > 0 {
		suffix = ":"
	}
	fmt.Fprintf(conn, "%d|%s%s", len(names), strings.Join(names, ":"), suffix)
}

func cat(conn net.Conn, filename string) {
	size := 1024 * 1024
	ctx := make([]byte, size)
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Fprint(conn, "0|")
	}
	n, _ := file.Read(ctx)
	fmt.Fprintf(conn, "%d|%s", size, ctx[:n])
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
END:
	for {
		op, args := cmd(conn)
		fmt.Println(op, args)
		switch op {
		case "ls":
			ls(conn)
		case "cat":
			cat(conn, args[0])
		case "quit":
			break END
		}
	}
	log.Printf("client close: %s", conn.RemoteAddr())
}

func main() {
	// 定义日志文件
	logfile, _ := os.OpenFile("fileserver.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	log.SetOutput(logfile)
	defer logfile.Close()

	var addr = "127.0.0.1:6666"
	// 建立监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("create listen err： %s\n", err)
	}
	defer listener.Close()
	log.Printf("listen on [%s]\n", addr)

	// 循环接收客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("connect err: %s\n", err)
		}
		log.Printf("%s connected...\n", conn.RemoteAddr())
		// 处理客户端连接
		go HandleConn(conn)
	}

}
