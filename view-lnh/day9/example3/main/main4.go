// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https  http2

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello2)
	http.ListenAndServeTLS(":8443", "/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day9/example3/main/server.crt", "/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day9/example3/main/server.key", nil)
}

func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!!!")
}
