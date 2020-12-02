package main

import (
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		println(path, info.Name(), info.Size(), info.ModTime().Format("2006-01-02 15:03:04"))
		return err
	})
}
