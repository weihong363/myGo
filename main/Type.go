package main

import "fmt"

func main() {
	//类型断言,提供了访问接口值底层具体值的方式。
	var i interface{} = "aaaa"
	j := i.(string)
	fmt.Println(j)
	//为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
	//像这样
	k, isOk := i.(string)
	fmt.Println(k, isOk)

	//类型选择 是一种按顺序从几个类型断言中选择分支的结构。其实就是switch，但是传的类型是接口（这样可以处理多种类型）
	do(11)
	do("wqeqw")
	do(true)

	//Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
	//可以通过实现String()方法来自定义输出的值（有点像java重写toString）
	o := Object{"张三", "eeeee"}
	fmt.Println(o)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Object struct {
	Name  string
	Value string
}

func (o Object) String() string {
	return fmt.Sprintf("name:%v   value:%v", o.Name, o.Value)
}

//嵌入（嵌套）类型，go没有继承，通过组合的方式，来复用已有的代码
//初始化时，先初始化内部再初始化外部
//如果内部类型实现了某个接口，外部类型也会被认为实现了该接口

type Combine struct {
	Key string
	//组合了Object结构体
	Object
}
