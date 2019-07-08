// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package encapsulation

type mitix struct {
	Name string
	age  int
	nomo string
}

func NewMitix(name string, age int) *mitix {
	return &mitix{Name: name, age: age}
}

func (this *mitix) GetAge() int {
	return (*this).age
}
