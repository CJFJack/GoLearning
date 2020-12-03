package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var GoFileChan chan string

// 遍历目录下go代码文件
func recursive() {
	path := "."
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		suffix := filepath.Ext(path)
		if suffix == ".go" {
			GoFileChan <- path
		}
		return nil
	})
}

// 统计文件中代码行数
func fileLines(path string) (int, error) {
	file, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	reader := bufio.NewReader(file)
	n := 0
	for {
		ctx, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return n, err
		}
		s := strings.TrimSpace(string(ctx))
		if len(s) > 0 && !strings.HasPrefix(s, "//") {
			n++
		}
	}
	return n, nil
}

func main() {
	fmt.Println(time.Now())
	var wg sync.WaitGroup
	var rowsTotal int
	var locker sync.Mutex

	// 初始化文件管道
	GoFileChan = make(chan string, 10)
	// 遍历go代码文件，写入管道
	wg.Add(1)
	go func() {
		recursive()
		// 关闭管道
		close(GoFileChan)
		wg.Done()
	}()

	// 读取管道中文件，统计代码行数
	for path := range GoFileChan {
		wg.Add(1)
		go func(path string) {
			rows, err := fileLines(path)
			if err == nil {
				locker.Lock()
				rowsTotal += rows
				locker.Unlock()
			}
			wg.Done()
		}(path)
	}

	//for path := range GoFileChan {
	//	rows, err := fileLines(path)
	//	if err == nil {
	//		locker.Lock()
	//		rowsTotal += rows
	//		locker.Unlock()
	//	}
	//}

	wg.Wait()

	fmt.Println(rowsTotal)
	fmt.Println(time.Now())

}
