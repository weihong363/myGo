package main

import "fmt"

//接口定义如下
type Human interface {
	//定义一个抽象方法
	gender()
}

type Man struct {
}

//实现输出男人的方法
func (m Man) gender() {
	fmt.Println("I am a man...")
}

type Woman struct {
}

//实现输出女人的方法
func (w Woman) gender() {
	fmt.Println("I am a woman...")
}

//接口隐式实现（示例来自go的官方文档）
type I interface {
	M()
}

type T struct {
	S string
}

//此方法标识T实现了接口I，无需显式声明，直接实现I的方法
func (t T) M() {
	fmt.Println(t.S)
}

//接口值。接口可以像其他值那样作为参数或者返回值
//接口值保存了一个具体底层类型的具体值。
//接口值调用方法时会执行其底层类型的同名方法。

//底层的接口值为nil时（比如下面的i2），即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。保存了 nil 具体值的接口其自身并不为 nil。
//在go中采用如下方法处理

func (t *T) Me() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

//空接口，指定了零个方法的接口值被称为 *空接口：*
//空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
//空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
type Empty interface{}

func main() {
	//接口
	var a Human
	a = new(Man)
	a.gender()
	a = new(Woman)
	a.gender()

	var i I = T{"implemet"}
	var i2 I
	i.M()
	//输出接口
	fmt.Println(i, i2)

	//nil 接口值既不保存值也不保存具体类型。
	//为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
	//i2.M()

	var e Empty = "testmodule"
	fmt.Println(e)
	var e1 Empty = 3
	fmt.Println(e1)
}
