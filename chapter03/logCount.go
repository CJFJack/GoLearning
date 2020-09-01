package main

import "fmt"

/*
1. 每个IP出现的次数
2. 每个url出现的次数
3. 每个IP在每个url中出现的次数
*/

func main() {
	log := [][4]string{
		{"1.1.1.1", "/index.html", "200", "10000"},
		{"1.1.1.2", "/index.html", "200", "1000"},
		{"1.1.1.1", "/index.html", "200", "10000"},
		{"1.1.1.3", "/home.html", "200", "10000"},
		{"1.1.1.3", "/abc.html", "200", "10000"},
	}

	count1 := map[string]int{}
	for _, s := range log {
		//fmt.Println(s)
		count1[s[0]] += 1
	}
	fmt.Println(count1)

	count2 := map[string]int{}
	for _, s := range log {
		count2[s[1]] += 1
	}
	fmt.Println(count2)

	count3 := make(map[string]map[string]int, 0)
	fmt.Println(count3)
	for _, s := range log {
		if _, ok := count3[s[1]]; ok {
			count3[s[1]][s[0]] += 1
		} else {
			count3[s[1]] = make(map[string]int)
			count3[s[1]][s[0]] += 1
		}
	}
	fmt.Println(count3)
}
