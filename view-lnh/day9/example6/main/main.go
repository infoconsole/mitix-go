// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"view-lnh/day9/example6/module"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:mitix@tcp(127.0.0.1:3306)/golang-db0")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := OpenDB()
	defer db.Close()
	r, err := db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "Hong.Lvhanng", "man", "hlh19850220@163.com")
	if err != nil {
		fmt.Println("insert into db err", err)
		return
	}
	personId, _ := r.LastInsertId()
	fmt.Println(personId)

	rows, _ := db.Query("SELECT `username`, `sex`,`email` FROM `person` WHERE id=?", personId)

	person := module.Person{}
	for rows.Next() {
		rows.Scan(&person.UserName, &person.Sex, &person.Email)
		fmt.Printf("id is %d, name is %v, sex is %v, email is %v", personId, person.UserName, person.Sex, person.Email)
	}
}
