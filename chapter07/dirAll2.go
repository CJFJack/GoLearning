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

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		fpath := path + "/" + fileInfo.Name()
		if fileInfo.IsDir() {
			fmt.Println("dir:", fpath)
			Dir(fpath)
		} else {
			fmt.Println("file:", fpath)
		}
	}
}

func main() {

	Dir("testdir")

}
