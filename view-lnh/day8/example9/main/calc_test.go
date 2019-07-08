// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "testing"

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add(2,4) error ,expect:%d,actual:%d", 6, r)
	}
	t.Logf("test add(2,4) succ")
}

func TestSub(t *testing.T)  {
	r:=sub(8, 3)
	if r!=5{
		t.Fatalf("sub(8,3) error ,expect:%d,actual:%d", 5, r)
	}
	t.Logf("test sub(8,3) succ")
}