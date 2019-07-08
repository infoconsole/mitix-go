// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 葡萄类

package fruit

import "fmt"

type grape struct {
	seedless bool
}

func (p *grape) Grow() {
	fmt.Println("grape is grow")
}

func (p *grape) Harvest() {
	fmt.Println("grape is harvest")
}

func (p *grape) Plant() {
	fmt.Println("grape is plant")
}

func (p *grape) isSeedless() bool {
	return p.seedless
}
