package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file1, _ := os.Open("test1.txt")
	file2, _ := os.Open("test2.txt")
	file3, _ := os.Open("multiReader.go")
	reader := io.MultiReader(file1, file2, file3)

	ctx := make([]byte, 5)
	for {
		n, err := reader.Read(ctx)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		fmt.Print(string(ctx[:n]))
	}

	defer file1.Close()
	defer file2.Close()
}
