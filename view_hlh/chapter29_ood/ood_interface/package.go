// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 典型的接口访问   显示了go语言的多肽

package ood_interface

//1) 接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的 多态和高内聚低偶合的思想。
//
//2) Golang 中的接口，不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个 变量就实现这个接口。因此，Golang 中没有 implement 这样的关键字

//1) 接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
//2) 接口中所有的方法都没有方法体,即都是没有实现的方法。
//
//3) 在 Golang 中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现 了该接口。
//
//4) 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型
//
//5) 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
