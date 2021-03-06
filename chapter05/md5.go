package main

import (
	"crypto/md5"
	"encoding/hex"
)
import "fmt"

func main() {
	// 方式一：
	fmt.Printf("%x\n", md5.Sum([]byte("我是Jack")))

	// 方式二：
	hasher := md5.New()
	hasher.Write([]byte("我是"))
	hasher.Write([]byte("Jack"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))

}
