// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 9.golang中的struct没有构造函数，一般可以使用工厂模式来解决这个问题

package main

import (
	"fmt"
	module2 "mitix-go/view_hlh/chapter29_ood/ood_factory/module"
)

func main() {
	s := module2.NewStudent("tony", 20)
	fmt.Println(s)
}
