package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("我是jack\n"))
	builder.WriteString("我是LULU")
	fmt.Println(builder.String())
	builder.Reset()
	fmt.Println(builder.String())
}
