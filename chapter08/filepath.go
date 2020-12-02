package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs(".")
	fmt.Println(path)
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Dir(path))

	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Split(path))

	paths := "path1;path2;path3"
	fmt.Printf("%T", filepath.SplitList(paths))
	fmt.Println(filepath.SplitList(paths))

	fmt.Println(filepath.Glob("./"))

	fmt.Println(filepath.Match("./*.go", "./filepath.go"))
	fmt.Println(filepath.Match("./*.text", "./filepath.go"))

}
