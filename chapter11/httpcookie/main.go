package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func ParseCookie(cookie string) map[string]string {
	cookieMap := make(map[string]string)
	if strings.TrimSpace(cookie) == "" {
		return cookieMap
	}
	values := strings.Split(cookie, ";")
	for _, value := range values {
		element := strings.Split(value, "=")
		cookieMap[strings.TrimSpace(element[0])] = strings.TrimSpace(element[1])
	}
	return cookieMap
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 方式一：
		// 解析cookie
		cookie := ParseCookie(request.Header.Get("Cookie"))
		// 读取cookie value
		counter := 0
		if v, err := strconv.Atoi(cookie["counter"]); err == nil {
			counter = v

		}

		// 方式二：
		// 解析并读取cookie value
		cookies, _ := request.Cookie("counter")
		fmt.Printf("%T, %#v\n", cookies.Value, cookies.Value)

		// 设置Cookie
		counterCookie := &http.Cookie{
			Name:     "counter",
			Value:    strconv.Itoa(counter + 1),
			HttpOnly: true,
		}
		http.SetCookie(writer, counterCookie)

	})

	http.ListenAndServe(":8888", nil)
}
