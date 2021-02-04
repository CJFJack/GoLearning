package service

import (
	"log"
	"rpcserver/data"
	"time"
)

// 定义计算服务结构体
type Calculator struct{}

// 定义Add方法
func (c *Calculator) Add(request *data.CalculatorRequest, response *data.CalculatorResponse) error {
	log.Println("[+] call add method")
	response.Result = request.Left + request.Right
	time.Sleep(time.Second * 5)
	return nil
}
