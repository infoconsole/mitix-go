// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package module

type Person struct {
	Id int `db:"id"`
	UserName string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}
