// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package balance

import (
	"net/url"
)

type RoundRobin struct {
	pos         int
	invokerList InvokerList
}

func (roundRobin *RoundRobin) AddInvokers(list InvokerList) {
	roundRobin.invokerList = list
}

func (roundRobin *RoundRobin) Select(url url.URL) Invoker {
	pos := roundRobin.pos
	invokers := roundRobin.invokerList.invokers
	invoker := invokers[pos]

	pos ++
	if pos == len(invokers)-1 {
		pos = 0
	}
	roundRobin.pos = pos
	return invoker
}
