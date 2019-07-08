// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	t, err := template.ParseFiles("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day9/example5/main/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", Age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
