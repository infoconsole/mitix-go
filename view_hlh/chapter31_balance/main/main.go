// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"mitix-go/view_hlh/chapter31_balance/balance"
	"net/url"
)

func main() {
	var invoker1 = balance.Invoker{
		Host:   "192.168.2.1:80",
		Weight: 20,
	}
	var invoker2 = balance.Invoker{
		Host:   "192.168.2.1:80",
		Weight: 15,
	}
	var invoker3 = balance.Invoker{
		Host:   "192.168.2.1:80",
		Weight: 30,
	}
	var invoker4 = balance.Invoker{
		Host:   "192.168.2.1:80",
		Weight: 100,
	}
	var invoker5 = balance.Invoker{
		Host:   "192.168.2.1:80",
		Weight: 50,
	}
	var invoker6 = balance.Invoker{
		Host:   "192.168.2.1:80",
		Weight: 10,
	}
	var invokerList balance.InvokerList
	invokerList.AddInvoker(invoker1).AddInvoker(invoker2).AddInvoker(invoker3).AddInvoker(invoker4).
		AddInvoker(invoker5).AddInvoker(invoker6)
	fmt.Println(invokerList)

	//var balance balance.LoadBalance = &balance.RoundRobin{}
	var balance balance.LoadBalance =&balance.WeightRoundRobin{}
	balance.AddInvokers(invokerList)
	var url = url.URL{
		Host: "ag.mitix.com",
		Path: "/info",
	}

	for i := 0; i < 200; i++ {
		dinvoker := balance.Select(url)
		fmt.Println(dinvoker)
	}
}
