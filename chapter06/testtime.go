package main

import (
	"fmt"
	"time"
)

var Local = time.Local

func main() {

	now := time.Now()
	fmt.Printf("%T, %#v\n", now, now)

	year := now.Year()
	month := now.Month()
	day := now.Day()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println(year, month, day, hour, minute, second)

	fmt.Println(time.Now().Unix())
	fmt.Printf("%T, %#v\n", time.Unix(1600846466, 0), time.Unix(1600846466, 0))
	fmt.Printf("%T, %#v\n", time.Unix(1600846466, 0).Format("2006-01-02 15:04:05"), time.Unix(1600846466, 0).Format("2006-01-02 15:04:05"))

	fmt.Println(time.Date(year, month, day, 0, 0, 0, 0, time.Local))

	ctime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-09-23 15:55:00", time.Local)
	fmt.Println(ctime)

	fmt.Println(time.Until(ctime))
	fmt.Println(time.Since(ctime))

	duration, _ := time.ParseDuration("1h1m1s")
	fmt.Println(duration.Hours())
	fmt.Println(now.Add(duration))

	duration, _ = time.ParseDuration("24h")
	yesterday := now.Add(duration)
	fmt.Println(now.Before(yesterday))
	fmt.Println(now.After(yesterday))

}
