package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("abcdefghijklmnopqrstuvwxyz")
	ctx := make([]byte, 10)
	for {
		n, err := reader.Read(ctx)
		if err == io.EOF {
			break
		}
		fmt.Println(string(ctx[:n]))
	}
}
