package main

import "fmt"

var a int = 1 //声明赋值同时进行
//a =bool //错误赋值
var b = 2 // 声明未加类型，自动推断类型
a, b = 3, 4 //多变量同时赋值，只能在函数体内
//c := true   //短类型声明赋值,只能在函数体内
func main() {
	a, b = 3, 4 //多变量同时赋值，只能在函数体内
	c := true   //短类型声明赋值,只能在函数体内

	fmt.Printf("a address: %v  value: %v \n", &a, a)
	fmt.Printf("b address: %v  value: %v \n", &b, b)
	fmt.Printf("c address: %v  value: %v \n", &c, c)

}
