package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Printf("%x\n", []byte("我是啦啦啦"))
	fmt.Println(hex.EncodeToString([]byte("我是啦啦啦")))
	txt, _ := hex.DecodeString("e68891e698afe595a6e595a6e595a6")
	fmt.Println(string(txt))

}
