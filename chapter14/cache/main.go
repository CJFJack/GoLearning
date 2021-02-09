package main

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

type User struct {
	ID   int
	Name string
}

func main() {
	//gob.Register(new(User))
	//config, _ := config.NewCache("memory", `{"interval":60}`)
	cache, err := cache.NewCache("file", `{"CachePath":"./config","FileSuffix":".config","EmbedExpiry":"60", "DirectoryLevel":"3"}`)
	//config, err := config.NewCache("redis", `{"key":"172.16.253.14","conn":":6379","dbNum":"0","password":""}`)
	fmt.Println("err", err)

	// 获取cache
	fmt.Println(cache.Get("name"))
	// 设置cache
	fmt.Println("put")
	cache.Put("name", "cjf", time.Minute)
	fmt.Println(cache.Get("name"))
	// 判断是否存在
	fmt.Println(cache.IsExist("name"))
	// 删除cache
	//fmt.Println("delete")
	//config.Delete("name")
	//fmt.Println(config.Get("name"))

	user := User{1, "cjf"}
	cache.Put("user", user, time.Minute)
	user2 := cache.Get("user")
	fmt.Println(user2.(User).Name)
}
