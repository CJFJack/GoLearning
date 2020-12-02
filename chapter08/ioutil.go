package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FormatSize(size int64) string {
	units := [6]string{"B", "KB", "MB", "GB", "TB", "GB"}
	index := 0
	fSize := float64(size)

	for fSize > 1024 && index < len(units)-1 {
		fSize /= 1024
		index++
	}
	return fmt.Sprintf("%.2f%s", fSize, units[index])
}

func main() {

	content, err := ioutil.ReadFile("filepath.go")
	fmt.Println(string(content), err)

	ioutil.WriteFile("test.txt", []byte("heiheihei"), os.ModePerm)

	files, err := ioutil.ReadDir(".")
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("D", file.Name(), file.ModTime().Format("2006-01-02 15:03:04"))
		} else {
			fmt.Println("F", file.Name(), FormatSize(file.Size()), file.ModTime().Format("2006-01-02 15:03:04"))
		}
	}

	//file, _ := os.Open("filepath.go")
	//reader := bufio.NewReader(file)
	//ctx, _ := ioutil.ReadAll(reader)
	//fmt.Println(string(ctx))

	files, _ = ioutil.ReadDir(".")
	for _, file := range files {
		fmt.Println(file.Name())
	}

}
