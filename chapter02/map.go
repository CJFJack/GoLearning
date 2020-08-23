package main

import (
	"fmt"
)

func main() {
	var scores map[string]float64
	fmt.Printf("%T, %#v\n", scores, scores)

	scores = map[string]float64{"1": 1, "2": 2, "3": 3}
	fmt.Printf("%T, %#v\n", scores, scores)

	v, ok := scores["1"]
	fmt.Println(v)
	fmt.Println(ok)

	// 修改
	scores["1"] = 98
	fmt.Printf("%T, %#v\n", scores, scores)

	// 删除
	delete(scores, "1")
	fmt.Printf("%T, %#v\n", scores, scores)

	// 遍历
	for k, v := range scores {
		fmt.Println(k, v)
	}

	//info := make(map[string]string)
	info := map[string]string{}
	fmt.Printf("%T, %#v\n", info, info)
	info["aa"] = "1"
	fmt.Printf("%T, %#v\n", info, info)

	article := "I have a dream"
	stat := map[rune]int{}

	for _, ch := range article {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			stat[ch]++
		}
	}

	for ch, cnt := range stat {
		fmt.Printf("%c: %d\n", ch, cnt)
	}

}
