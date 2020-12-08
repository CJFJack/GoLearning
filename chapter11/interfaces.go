package main

import (
	"fmt"
	"net"
)

func main() {

	intfs, _ := net.Interfaces()

	for _, intf := range intfs {
		fmt.Println(intf.Name, intf.HardwareAddr, intf.MTU)
		addrs, _ := intf.Addrs()
		for _, addr := range addrs {
			fmt.Println(intf.Name, addr)
		}
	}

}
