package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	os.Stdout.WriteString("hello")

	fmt.Println(os.TempDir())
	fmt.Println(os.UserHomeDir())

	fmt.Println(os.UserCacheDir())

	path, _ := os.Executable()
	fmt.Println(path)
	fmt.Println(filepath.Dir(path))

}
