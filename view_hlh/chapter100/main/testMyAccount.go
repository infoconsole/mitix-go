// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	viewMenu()
	var chose int8
	fmt.Scanf("%d", &chose)
	switch chose {
	case 1:
		viewContent(chose)

	}
	
}

func viewMenu() {
	fmt.Println("-----------------家庭收支记账软件-----------------")
	fmt.Println("                 1.收支明细")
	fmt.Println("                 2.登记收入")
	fmt.Println("                 3.登记支出")
	fmt.Println("                 4.退    出")
	fmt.Println("")
	fmt.Print("                 请选择（1-4）:")
}
