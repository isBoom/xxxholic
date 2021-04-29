package main

import (
	"xxxholic/conf"
	"xxxholic/server"
)

func main() {
	conf.Init()
	r := server.NewRouter()
	r.Run(":8000")
}
