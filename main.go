package main

import (
	"xxxholic/conf"
	"xxxholic/server"
)

func main() {
	conf.Init()
	// 装载路由
	r := server.NewRouter()
	r.Run(":8000")

}
