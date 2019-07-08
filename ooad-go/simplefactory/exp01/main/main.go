// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "mitix-go/ooad-go/simplefactory/exp01/fruit"

func main() {
	fruiter := fruit.NewFruit("apple")
	fruiter.Grow()
	fruiter.Harvest()
	fruiter.Plant()


	fruiter= fruit.NewFruit("strawberry")
	fruiter.Grow()
	fruiter.Harvest()
	fruiter.Plant()
}
