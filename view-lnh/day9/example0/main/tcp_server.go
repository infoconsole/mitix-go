// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// tcp服务端

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting socket server")
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Println("listen failed,error ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err", err)
			return
		}
		fmt.Println("read : ", string(buf[:n]))
	}
}
