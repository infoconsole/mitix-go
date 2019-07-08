// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package fruit

func NewFruit(t string) (fruit Fruiter) {
	switch t {
	case "apple":
		fruit = &apple{}
	case "grape":
		fruit = &grape{}
	case "strawberry":
		fruit = &strawberry{}
	}
	return fruit
}
