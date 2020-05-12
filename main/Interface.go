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

func main() {
	//接口
	var a Human
	a = new(Man)
	a.gender()
	a = new(Woman)
	a.gender()
}
