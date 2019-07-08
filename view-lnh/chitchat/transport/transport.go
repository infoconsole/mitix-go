// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package transport

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"view-lnh/chitchat/proto"
)

func ReadPackage(conn net.Conn) (msg proto.Message, err error) {
	//8k buffer
	var buf [8192]byte
	n, err := conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("log: read header failed")
		return
	}
	fmt.Println("log: read package:", buf[0:4])

	var packLen uint32
	packLen = binary.BigEndian.Uint32(buf[0:4])

	fmt.Printf("log: receive len:%d", packLen)
	n, err = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("log: read body failed")
		return
	}

	fmt.Printf("log: receive data:%s\n", string(buf[0:packLen]))
	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("log: unmarshal failed, err:", err)
	}
	return
}

func WritePackage(conn net.Conn, data []byte) (err error) {
	var buf [8192]byte
	packLen := uint32(len(data))
	//write header
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	n, err := conn.Write(buf[0:4])

	if err != nil {
		fmt.Println("log: write header failed")
		return
	}

	n, err = conn.Write(data)
	if err != nil {
		fmt.Println("log: write data  failed")
		return
	}

	if n != int(packLen) {
		err = errors.New("log: write data not fninshed")
		return
	}

	return
}
