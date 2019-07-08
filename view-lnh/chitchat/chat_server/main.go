package main

import (
	"fmt"
	"time"
	"view-lnh/chitchat/chat_server/server"
	"view-lnh/chitchat/model"
)

func main() {
	p, err := server.InitRedis("127.0.0.1:6379", 16, 1024, time.Second*300)
	if err != nil {
		fmt.Println(err)
	}
	server.AppContext.Pool = p
	server.AppContext.UserMgr = model.NewUserMgr(p)
	server.StartServer("0.0.0.0", 8090)
}
