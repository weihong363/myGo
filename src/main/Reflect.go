package main

import (
	"fmt"
	"reflect"
)

func main() {
	//反射，就是建立在类型之上的，Golang的指定类型的变量的类型是静态的（也就是指定int、string这些的变量，它的type是static type），
	//在创建变量的时候就已经确定，反射主要与Golang的interface类型相关（它的type是concrete type），只有interface类型才有反射一说。
	//在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型：(value, type)
	//value是实际变量值，type是实际变量的类型。一个interface{}类型的变量包含了2个指针，一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。
	//golang的反射包给我们提供了TypeOf()和ValueOf()两个方法来获取type和value
	//示例如下：
	i := int(8)
	//从接口类型变量获取反射类型变量
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))

	//通过反射来转换类型
	//1.转换的时候，如果转换的类型不完全符合，则直接panic，类型要求非常严格！
	//2.转换的时候，要区分是指针还是指
	//3.也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
	//如：
	var f float32 = 3.1415926
	//typeOf := reflect.TypeOf(f)
	valueOf := reflect.ValueOf(f)
	//获取类型
	t := valueOf.Type()
	//类型只能填float32，用其他类型会报错
	elem := valueOf.Interface().(float32)
	//输出时会丢失精度
	fmt.Println(elem)
	fmt.Println(t)
	//即使是int转string也是不行的
	//s := reflect.ValueOf(i).Interface().(string)
	//fmt.Println(s)

	//通过反射获取结构体内的所有字段和接口（如果没初始化，取出来就是空的，有初始化取出来才有值）
	tt := TestType{"2133", "421321", "A", "B"}
	of := reflect.ValueOf(tt)
	for i := 0; i < of.NumField(); i++ {
		fmt.Println(of.Field(i))
	}

	//这个输出的是方法的地址。。。有个Name属性一直拿不到
	//todo 反射获取结和字段部分后续还要再看下
	for i := 0; i < of.NumMethod(); i++ {
		fmt.Println(of.Method(i))
	}

	//修改字段的值
	value := reflect.ValueOf(&tt.A)
	value.Elem().SetString("hahahaha")
	//现在可以看到，tt的A的值已经变了
	fmt.Println(tt)

	//动态调用方法
	//获取方法
	method := of.MethodByName("TestMethod")
	//方法的参数
	args := []reflect.Value{reflect.ValueOf("动态调用方法！")}
	//调用
	method.Call(args)
}

type TestType struct {
	A string
	B string
	InterfaceA
	InterfaceB
}
type InterfaceA interface {
}
type InterfaceB interface {
}

func (tt TestType) TestMethod(input string) {
	fmt.Println(input)
}
