package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	//hello world
	fmt.Println("hello world!")

	//变量定义(var 参数名，参数2... 类型 = 值,值2...)  没被使用的变量会标红
	var a string = ""
	//var b,c int  = 1,2
	//此外，还能这么写
	//var (
	//	d bool = true
	//)
	var e = 1
	//var f int
	//等同于 var g string = ""
	//g := ""

	//常量定义
	const PI float32 = 3.1415926
	const (
		G float32 = 9.231242432
	)
	/**
	 * iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	 * 可以用来做枚举
	 */
	//如：这里y,z分别为1，2
	const (
		x = iota
		y
		z
	)
	fmt.Print(a)

	//运算符，加减乘除求模取余，还有关系运算符比较大小的，逻辑运算符，位运算符，赋值运算符这些和java一样，就不说了
	//有两个不同的运算符 &和*，&获取变量地址,*表示指针变量

	var ptr *int
	//指针指向e的地址。指针变量 * 和地址值 & 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。
	//当变量前面有 * 标识时，取的是该地址里的值，否则会直接输出地址值。
	ptr = &e
	fmt.Println(&e)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
