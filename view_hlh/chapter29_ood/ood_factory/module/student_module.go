// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package module

type student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) *student {
	return &student{
		Name: name,
		Age:  age,
	}
}
