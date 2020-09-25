package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(math.Abs(-1))
	fmt.Println(math.Ceil(1.49))
	fmt.Println(math.Floor(-1.1))
	fmt.Println(math.Max(-1, 5))
	fmt.Println(math.Min(-1, 5))

	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(101))
	fmt.Println(rand.Float32())
}
