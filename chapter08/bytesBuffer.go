package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBuffer([]byte("lulu\n"))

	buffer.Write([]byte("cjf"))

	ctx, _ := buffer.ReadBytes('j')
	fmt.Println(string(ctx))
	buffer.String()
}
