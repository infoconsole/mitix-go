// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package balance

import "net/url"

type Invoker struct {
	Host       string
	Weight     int
	RealWeight int
}

type InvokerList struct {
	size     int
	invokers []Invoker
}

func (invokers *InvokerList) Init() {
	invokers.size = 0
	invokers.invokers = make([]Invoker, 0)

}

func (this *InvokerList) AddInvoker(invoker Invoker) *InvokerList {
	this.size++
	this.invokers = append(this.invokers, invoker)
	return this
}

type LoadBalance interface {
	Select(url url.URL) Invoker
	AddInvokers(invokers InvokerList)
}
