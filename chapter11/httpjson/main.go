package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//io.Copy(os.Stdout, request.Body)
		var info map[string]interface{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&info)
		fmt.Println(info)
	})

	http.ListenAndServe(":8888", nil)
}
