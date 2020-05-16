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
	//匿名函数
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

//也可以为非结构体类型声明方法。如：

type F float32

func (f F) M() float32 {
	return 1.00
}

//只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
//接收者(上面例子里的f就是接收者)的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。
//接收者可以为指针，但是不能为*int这样的（毕竟要是同一个包内的）
//使用指针接收者的原因有二：
//首先，方法能够修改其接收者指向的值。
//其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
//以下为摘自官网的示例
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func mainTwo() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
