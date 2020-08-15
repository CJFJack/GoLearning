/*
生成随机数字
提示用户在控制台输入猜测的数字
比较，当用户输入较大，提示太大了
当用户输入太小，提示太小了
当用户输入正确，提示经过N次对了，太聪明了
用户最多猜5次，如果5出内没有正确，提示太笨了，游戏结束

扩展，当成功或者失败后，提示用户是否继续，输入yes，y,Y则继续，重新生成随机数
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

BEGIN:
	var (
		random int
		input  int
		n      int
		min    = 0
		max    = 100
		N      = 5
		gohead = "no"
	)
	rand.Seed(time.Now().Unix())
	random = rand.Intn(max)
	fmt.Println("请输入需要猜测的数字：")
	_, _ = fmt.Scan(&input)
	n++
	for {
		if input == random {
			fmt.Printf("经过%d次猜对了，太聪明了\n", n)
			fmt.Println("是否要继续玩，输入yes或y，则继续，否则退出")
			_, _ = fmt.Scan(&gohead)
			if gohead == "yes" || gohead == "y" {
				goto BEGIN
			} else {
				break
			}
		} else if input > random {
			max = input
			fmt.Printf("太大了，请重新输入(%d ~ %d)：\n", min, max)
			_, _ = fmt.Scan(&input)
			n++
		} else if input < random {
			min = input
			fmt.Printf("太小了，请重新输入(%d ~ %d)：\n", min, max)
			_, _ = fmt.Scan(&input)
			n++
		}
		if n > N-1 {
			fmt.Println("太笨了")
			fmt.Println("是否要继续玩，输入yes、y或Y，则继续，否则退出")
			_, _ = fmt.Scan(&gohead)
			if gohead == "yes" || gohead == "y" || gohead == "Y" {
				goto BEGIN
			} else {
				break
			}
		}
	}

}
