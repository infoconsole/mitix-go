// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package server

import (
	"github.com/gomodule/redigo/redis"
	"view-lnh/chitchat/model"
)

var AppContext = &ApplicationContext{}

type ApplicationContext struct {
	Pool    *redis.Pool
	UserMgr *model.UserMgr
}
