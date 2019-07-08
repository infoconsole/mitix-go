// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package balance

import (
	"math"
	"net/url"
)

type WeightRoundRobin struct {
	pos         int
	wpos        int
	invokerList InvokerList
}

func (wrobin *WeightRoundRobin) AddInvokers(list InvokerList) {
	var mWeight = list.invokers[0].Weight
	var coefficient = 1
	for _, invoker := range list.invokers {
		invoker.RealWeight = invoker.Weight
		if invoker.Weight < mWeight {
			mWeight = invoker.Weight
		}
	}
	if mWeight < 2 {
		return
	}
	for ; ; {
		coefficient = coefficient * 2
		coe := mWeight / coefficient
		if coe > 1/2 && coe < 2 {
			break
		}
	}
	for i := 0; i < list.size; i++ {
		var rweight = math.Ceil(float64(list.invokers[i].Weight / coefficient))
		list.invokers[i].RealWeight = int(rweight)
	}
	wrobin.invokerList = list
}

func (wrobin *WeightRoundRobin) Select(url url.URL) Invoker {
	pos := wrobin.pos
	wpos := wrobin.wpos
	invokers := wrobin.invokerList.invokers
	invoker := invokers[pos]
	wpos ++
	if wpos == invoker.RealWeight {
		pos ++
		wpos = 0
	}
	if pos == len(invokers)-1 {
		pos = 0
	}
	wrobin.pos = pos
	wrobin.wpos = wpos
	return invoker
}
