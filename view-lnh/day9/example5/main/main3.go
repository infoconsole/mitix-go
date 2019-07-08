// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	temp *template.Template
)

type IndexHandler struct{}

type WorldHandler struct{}

type Man struct {
	Name string
	Age  int
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := Man{Name: "Mary", Age: 30}
	if err := temp.Execute(w, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world !!!")
}

func initTemplate() {
	t, err := template.ParseFiles("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day9/example5/main/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	temp = t
}

func main() {
	initTemplate()
	indexHandler := IndexHandler{}
	worldHandler := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/", &indexHandler)
	http.Handle("/world", &worldHandler)

	server.ListenAndServe()

}
