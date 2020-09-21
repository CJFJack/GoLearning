package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 标准base64 使用 0-9 a-z A-Z + _ 共64个字符，随机生成的加密字符串
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("I am Jack")))
	txt, _ := base64.StdEncoding.DecodeString("SSBhbSBKYWNr")
	fmt.Println(string(txt))
	// 标准base64Url + => -  / => _
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("https://cmdb.61info.cn")))
	txt, _ = base64.URLEncoding.DecodeString("aHR0cHM6Ly9jbWRiLjYxaW5mby5jbg==")
	fmt.Println(string(txt))

	fmt.Printf("\n")

	// 非标准 3的倍数
	fmt.Println(base64.RawStdEncoding.EncodeToString([]byte("I am Jack")))
	txt, _ = base64.RawStdEncoding.DecodeString("SSBhbSBKYWNr")
	fmt.Println(string(txt))
	// 非标准base64Url 3的倍数  + => -  / => _
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("https://cmdb.61info.cn")))
	txt, _ = base64.RawURLEncoding.DecodeString("aHR0cHM6Ly9jbWRiLjYxaW5mby5jbg")
	fmt.Println(string(txt))

}
