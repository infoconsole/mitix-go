package server

import (
	"fmt"
	"net"
	"strconv"
)

func StartServer(ip string, port int) (err error) {
	fmt.Println("starting socket server")
	addr := ip + ":" + strconv.Itoa(port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen failed, ", err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept failed, ", err)
			continue
		}
		go proccess(conn)
	}
}

func proccess(conn net.Conn) {
	//返回执行错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("client process error", err)
		}
	}()
	defer conn.Close()
	client := &Client{
		conn: conn,
	}

	err := client.Process()
	if err != nil {
		fmt.Println("client process failed, ", err)
		return
	}
}
