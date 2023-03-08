package main

import "fmt"

const Pi float32 = 3.1415926
const e = 2.71828 //可以忽略类型

//多个常量同时定义，常用于枚举
const (
	Red    = 0
	Yellow = 1
	Blue   = 2
)

//用iota来生成枚举值
const (
	a0 = iota
	a1 = iota
	a2 = iota
)

//用iota来快速生成枚举值
const (
	b0 = iota
	b1
	b2
)

//用iota来快速生成枚举值
const (
	c0 = iota
	c1
	c2
	c3 = 5
	c4 = iota
	c5
)

func main() {
	fmt.Printf("%v,%v,%v \n", a0, a1, a2)
	fmt.Printf("%v,%v,%v \n", b0, b1, b2)
	fmt.Printf("%v,%v,%v,%v,%v,%v  \n", c0, c1, c2, c3, c4, c5)

}
