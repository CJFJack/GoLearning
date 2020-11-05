package main

import (
	"fmt"
	"os"
)

func Dir(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	names, err := file.Readdirnames(-1)
	if err != nil {
		return
	}

	for _, name := range names {
		fpath := path + "/" + name
		if fileInfo, err := os.Stat(fpath); err == nil {
			if fileInfo.IsDir() {
				fmt.Println("dir:", fpath)
				Dir(fpath)
			} else {
				fmt.Println("file:", fpath)
			}
		}
	}
}

func main() {

	Dir("testdir")

}
