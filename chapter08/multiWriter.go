package main

import (
	"io"
	"os"
)

func main() {
	file1, _ := os.Create("test1.txt")
	file2, _ := os.Create("test2.txt")
	writer := io.MultiWriter(file1, file2, os.Stdout)
	writer.Write([]byte("Hello World!!!"))
	defer file1.Close()
	defer file2.Close()
}
