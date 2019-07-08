// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 定义水果的一些特性  是一个接口

package fruit

type Fruiter interface {
	Grow()

	Harvest()

	Plant()
}
