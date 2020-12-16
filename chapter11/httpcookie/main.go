package main

import (
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
		// 解析cookie
		cookie := ParseCookie(request.Header.Get("Cookie"))

		// 设置Cookie
		counter := 0
		if v, err := strconv.Atoi(cookie["counter"]); err == nil {
			counter = v

		}
		counterCookie := &http.Cookie{
			Name:     "counter",
			Value:    strconv.Itoa(counter + 1),
			HttpOnly: true,
		}
		http.SetCookie(writer, counterCookie)

	})

	http.ListenAndServe(":8888", nil)
}
