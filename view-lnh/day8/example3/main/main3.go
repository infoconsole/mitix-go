// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

type student struct {
	name string
}

func main() {
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)

	stu1 := student{
		name: "oldflame-Jm",
	}

	stuChan <- &stu1
	stuOut := <-stuChan
	stu2, ok := stuOut.(student)
	if !ok {
		fmt.Println("can not convert", ok)
		return
	}
	fmt.Println(stu2)

}
