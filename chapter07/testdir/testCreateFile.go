package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	os.Create("testFile.csv")

	file, err := os.Open("testdir/name.txt")
	fmt.Println(err)
	defer file.Close()

	ctx := make([]byte, 10)
	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		fmt.Println(string(ctx[:n]))
	}

}
