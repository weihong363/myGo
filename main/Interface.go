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

func main() {
	//接口
	var a Human
	a = new(Man)
	a.gender()
	a = new(Woman)
	a.gender()

	var i I = T{"implemet"}
	i.M()
}
