package main

import (
	"fmt"
	"sort"
)

func main() {
	// char => counter
	// A=>65
	// B=>66
	// C=>67
	// D=>68
	state := [][]int{{'A', 3}, {'B', 2}, {'C', 1}, {'D', 2}}
	sort.Slice(state, func(i, j int) bool {
		return state[i][1] < state[j][1]
	})

	sort.SliceStable(state, func(i, j int) bool {
		return state[i][1] < state[j][1]
	})
	fmt.Println(state)

	index := sort.Search(len(state), func(i int) bool {
		return state[i][1] > 1
	})
	fmt.Println(index)

}
