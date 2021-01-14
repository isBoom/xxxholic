package main

import (
	"fmt"
	"xxxholic/conf"
	"xxxholic/server"
)

func main(){
	fmt.Println("go xxxholic")

	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":8000")

}
