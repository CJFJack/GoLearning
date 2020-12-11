package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:9999"
	packageListen, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatalf("create listen err: %s", err)
	}
	defer packageListen.Close()
	log.Printf("listen on [%s]", addr)

	for {
		// 接收消息
		ctx := make([]byte, 1024)
		n, remoteAddr, err := packageListen.ReadFrom(ctx)
		if err != nil {
			log.Printf("read from %s err: %s", remoteAddr, err)
		}
		log.Printf("read from %s: %s", remoteAddr, string(ctx[:n]))
		// 回复消息
		n, err = packageListen.WriteTo([]byte("received msg ok"), remoteAddr)
		fmt.Println(n, err)
	}
}
