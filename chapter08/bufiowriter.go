package main

import (
	"bufio"
	"os"
)

func main() {

	file, _ := os.Create("bufiowrite.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.Write([]byte("adbasdfasf\n"))
	writer.WriteString("123421153454")
}
