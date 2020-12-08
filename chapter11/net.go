package main

import (
	"fmt"
	"net"
)

func main() {
	// ip、端口、域名的常用方法
	fmt.Println(net.JoinHostPort("127.0.0.1", "6666"))
	fmt.Println(net.JoinHostPort("::1", "6666"))
	fmt.Println(net.SplitHostPort("127.0.0.1:9999"))
	fmt.Println(net.SplitHostPort("[::1]:9999"))

	fmt.Println(net.LookupAddr("127.0.0.1"))
	fmt.Println(net.LookupHost("ZHANGHONGQIANG.ad61v1.com."))
	fmt.Println(net.LookupHost("www.baidu.com"))
	fmt.Println(net.LookupIP("ZHANGHONGQIANG.ad61v1.com."))
	fmt.Println(net.LookupIP("www.baidu.com"))

	// ip结构体
	ip := net.ParseIP("127.0.0.1")
	fmt.Printf("%T, %#v\n", ip, ip)
	ip = net.ParseIP("::1")
	fmt.Printf("%T, %#v\n", ip, ip)

	// ipNet cidr格式
	ip, ipNet, err := net.ParseCIDR("127.0.0.1/24")
	fmt.Println(ip, ipNet, err)

	// ipNet是否包含ip，参数：ip结构体
	fmt.Println(ipNet.Contains(net.ParseIP("127.0.0.11")))
	fmt.Println(ipNet.Contains(net.ParseIP("192.168.1.1")))
}
