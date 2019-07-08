// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"view-lnh/day9/example6/module"
)

func main() {
	dbx, err := sqlx.Open("mysql", "root:mitix@tcp(127.0.0.1:3306)/golang-db0")
	if err != nil {
		panic(err)
	}
	defer dbx.Close()

	r, err := dbx.Exec("insert into person(username, sex, email)values(?, ?, ?)", "Hong.Lvhanng", "man", "hlh19850220@163.com")
	if err != nil {
		fmt.Println("insert person err", err)
		return
	}
	personId, _ := r.LastInsertId()
	fmt.Println(personId)

	var person []module.Person
	err = dbx.Select(&person, "SELECT `username`, `sex`,`email` FROM `person` WHERE id=?", personId)
	if err != nil {
		fmt.Println("select person err", err)
		return
	}
	fmt.Printf("id is %d, name is %v, sex is %v, email is %v", personId, person[0].UserName, person[0].Sex, person[0].Email)
}
