package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 打开目标文件
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 读取内容并写入内容
	//ctx := make([]byte, 10)
	//for {
	//	n, err := srcFile.Read(ctx)
	//	dstFile.Write(ctx[:n])
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//		return err
	//	}
	//}

	//fmt.Println(io.Copy(dstFile, srcFile))

	//buffer := make([]byte, 1024)
	//io.CopyBuffer(dstFile, srcFile, buffer)

	io.CopyN(dstFile, srcFile, 3)

	return err
}

func main() {
	src := "./test.txt"
	dst := "./test_copy.txt"
	err := copyFile(src, dst)
	if err != nil {
		fmt.Println(err)
	}

}
