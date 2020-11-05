package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// 加密
	// 方式一
	fmt.Printf("%x\n", []byte("我是啦啦啦"))
	// 方式二
	fmt.Println(hex.EncodeToString([]byte("我是啦啦啦")))

	// 解码
	txt, _ := hex.DecodeString("e68891e698afe595a6e595a6e595a6")
	fmt.Println(string(txt))

}
