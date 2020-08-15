package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{3, 6, 8, 9, 1, 3, 5, 4, 2, 1}
	sort.Ints(nums)
	fmt.Println(nums)
	//sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	//fmt.Println(nums)

	// search
	fmt.Println(sort.SearchInts(nums, 4))

	var names = []string{"abc", "l", "r", "k", "qwe"}
	sort.Strings(names)
	fmt.Println(names)
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println(names)
}
