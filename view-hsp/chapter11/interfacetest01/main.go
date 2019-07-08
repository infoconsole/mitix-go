// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package interfacetest01

import "fmt"

type Afac interface {
	test1(a int)
	test2()
}

type Bfac interface {
	test1(a, b int)
	test3()
}

type Student struct {
	name string
	age  int
}

func (this *Student) test1(a int) {
	fmt.Printf("this is a is %d", a)
}

func (this *Student) test1(a, b int) {
	fmt.Printf("this is a is %d,b is %d", a, b)
}
