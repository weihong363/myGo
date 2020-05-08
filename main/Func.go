package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(method("hello ", "world"))
	fmt.Println(swap("big", "small"))
	//骚操作，函数可以作为另外的函数的实参
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
	a := anonymous()
	fmt.Println(a)
}

//函数的定义如下：(默认情况下是值传递，不会影响到实际的参数)
//func function_name( [parameter list] ) [return_types] {
//   函数体
//}

//传入两个String参数，返回String
func method(a, b string) string {
	return a + b
}

//go函数可以返回多个值,在返回类型上要标好
func swap(a, b string) (string, string) {
	return b, a
}

//闭包，匿名函数
func anonymous() func() int {
	i := 1
	fmt.Println(i)
	r := func() int {
		return 5
	}
	return r
}

//Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
//所有给定类型的方法属于该类型的方法集。语法格式如下：
//func (variable_name variable_data_type) function_name() [return_type]{
//   /* 函数体*/
//}
type Circle struct {
	PI float32
}

func (c Circle) area(r float32) float32 {
	//PI是结构体里的属性
	return c.PI * r * r
}
