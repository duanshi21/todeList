package main

import (
	"fmt"
	"todoList/conf"
	"todoList/pkg/utils"
	"todoList/repository/db/dao"
	"todoList/router"
)

func main() {
	loading()
	r := router.NewRouter()
	fmt.Println("启动成功")
	_ = r.Run(conf.HttpPort)
}

func loading() {
	conf.Init()
	dao.MySqlInit() // 初始化MySQL连接
	utils.InitLog() // 初始化日志
}
