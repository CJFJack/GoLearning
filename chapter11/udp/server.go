package main

import (
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
		ctx := make([]byte, 1024)
		n, remoteAddr, err := packageListen.ReadFrom(ctx)
		if err != nil {
			log.Printf("read from %s err: %s", remoteAddr, err)
		}
		log.Printf("read from %s: %s", remoteAddr, string(ctx[:n]))
	}
}
