// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 苹果类

package fruit

import "fmt"

//create by factory
type apple struct {
	name    string
	treeAge int
}

func (p *apple) Grow() {
	fmt.Println("apple is grow")
}

func (p *apple) Harvest() {
	fmt.Println("apple is harvest")
}

func (p *apple) Plant() {
	fmt.Println("apple is plant")
}

func (p *apple) GetAge() int {
	return p.treeAge
}

func (p *apple) GetName() string {
	return p.name
}
