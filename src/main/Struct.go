package main

import "fmt"

//Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
//结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
//所以其实结构体可理解成java里的实体类

//结构体定义格式
//type struct_variable_type struct {
//   member definition
//   member definition
//   ...
//   member definition
//}

type Human struct {
	gender string
	age    int
}

func main() {
	//定义了结构体可以直接用于变量的声明
	zhangsan := Human{"男", 18}
	lisi := Human{"女", 16}

	fmt.Println(zhangsan, zhangsan.gender, zhangsan.age)
	fmt.Println(lisi, lisi.gender, lisi.age)

	//还有什么作为函数参数啥的，就没必要管了。。。把它当实体类理解就行

	//=======================================================

	//结构体指针——指定结构体指针变量，指向结构体的地址
	var ptr *Human
	ptr = &zhangsan
	fmt.Println(ptr, *ptr)
	//这不对啊，明明用&取了zhangsan的地址，为啥输出出来的会是&{男 18} 呢。。。，看下&张三是啥
	fmt.Println(&zhangsan, &ptr)
	//还是那样，查了下，是说fmt.Println(user, &user) 使用的是默认格式的打印方式%v。
	//对于 struct 默认输出格式是 {field0 field1 ...}。
	//想要输出地址，需要使用 %p，fmt.Printf("%v %p\n", user, &user)。
	//具体值要怎么打印，看官方文档https://golang.google.cn/pkg/fmt/#hdr-Printing
	//%v是打印结构体的,%p是打印地址的
	fmt.Printf("%v %p \n", zhangsan, &zhangsan)
}
