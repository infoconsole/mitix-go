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
	temp2 *template.Template
)

type IndexHandler2 struct{}

type WorldHandler2 struct{}

func (h *IndexHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//p := Man{Name: "Mary", Age: 30}
	//p也可以用map进行代替  不过需要小写
	//一般情况  map小写   模板的内容也是小写
	p := make(map[string]interface{})
	p["name"] = "honglvhang"
	p["age"] = 20
	//p2 := make(map[string]interface{})
	//p2["name"] = "honglvhang"
	//p2["age"] = 20
	//p3 := make(map[string]interface{})
	//p3["name"] = "honglvhang"
	//p3["age"] = 20
	//if err := temp2.Execute(w, []map[string]interface{}{p, p2, p3}); err != nil {
	if err := temp2.Execute(w, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

func (h *WorldHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world !!!")
}

func initTemplate2() {
	//t, err := template.ParseFiles("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day9/example5/main/index4.html")
	t, err := template.ParseFiles("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day9/example5/main/index5.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	temp2 = t
}

func main() {
	initTemplate2()
	indexHandler := IndexHandler2{}
	worldHandler := WorldHandler2{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/", &indexHandler)
	http.Handle("/world", &worldHandler)

	server.ListenAndServe()

}
