// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package vehicle

var carInstance *car

type car struct {
	name  string
	speed int
}

func NewCar() *car {
	if carInstance == nil {
		carInstance = &car{
			name:  "dika",
			speed: 100,
		}
	}
	return carInstance
}
