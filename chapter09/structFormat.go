package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type RedisConfig struct {
	IP   string
	PORT string
	AUTH int
	PASS string
}

type DbConfig struct {
	Host   string
	Port   int
	Uid    string
	Pwd    string
	DbName string
}

//Config 游戏服务器的配置
type Config struct {
	ServerId int
	Port     int //端口号

	Redis     *RedisConfig         `json:"redis" bson:"redis"`
	DbConfigs map[string]*DbConfig //如果配置多个数据库源，则用逗号分隔源的名字
	callbacks []func()
}

func (conf *Config) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}

func main() {

	conf := Config{
		ServerId: 1,
		Port:     8080,
		Redis:    &RedisConfig{"127.0.0.1", "6379", 1, "123"},
		DbConfigs: map[string]*DbConfig{
			"maindb": &DbConfig{
				Host: "127.0.0.1",
			},
		},
	}
	fmt.Println("Config:", conf.String())

}
