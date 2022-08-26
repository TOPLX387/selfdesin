package main

import (
	"fmt"
	"onepiece/database"
	"onepiece/packages/gredis"
	"onepiece/packages/setting"
	"onepiece/routers"
)

//数据库初始化，gedis初始化
func init() {
	setting.Setup()
	database.Setup()
	gredis.Setup()
}

//主函数
func main() {
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.Run(endPoint) //开启监听
}
