// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"view-lnh/chitchat/log"
	"view-lnh/chitchat/proto"
	"view-lnh/chitchat/transport"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("log: Error dialing", err.Error())
		return
	}
	defer conn.Close()
	log.ClientLog("Terminal is start ,please input command")
	for {
		var input string
		fmt.Scanf("%s\n", &input)
		if strings.HasPrefix(input, "log:") {
			continue
		}
		log.ClientLog(fmt.Sprintf("log: command [%s]\n", input))
		transport.WritePackage(conn, []byte(input))
		msg, err := transport.ReadPackage(conn)
		if err != nil {
			log.ClientLog(fmt.Sprintf("log: resp [%s]\n", msg))
		}
	}
}

func Cmdt() {
	var loginCmd proto.LoginCmd
	loginCmd.Id = 101
	loginCmd.Passwd = "123456789"

	data, err := json.Marshal(loginCmd)
	if err != nil {
		return
	}

	message := proto.Message{
		Cmd:  proto.UserRegister,
		Data: string(data),
	}

	msg, err := json.Marshal(message)
	if err != nil {
		return
	}
	fmt.Println(string(msg))
}
