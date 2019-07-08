// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package fruit

import "fmt"

type strawberry struct {
}

func (p *strawberry) Grow() {
	fmt.Println("strawberry is grow")
}

func (p *strawberry) Harvest() {
	fmt.Println("strawberry is harvest")
}

func (p *strawberry) Plant() {
	fmt.Println("strawberry is plant")
}
