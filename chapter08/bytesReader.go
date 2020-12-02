package main

import (
	"bytes"
	"fmt"
)

func main() {
	reader := bytes.NewReader([]byte("123123123123"))
	ctx := make([]byte, 3)
	n, _ := reader.Read(ctx)
	fmt.Println(string(ctx)[:n])

}
