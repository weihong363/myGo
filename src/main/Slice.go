package main

import "fmt"

func main() {
	//切片，slice是对数组的抽象。
	//Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
	//与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
	//声明一个不指定大小的数组来定义切片
	var s []int
	//也可以通过make()函数来创建切片,指定初始长度
	var s1 []int = make([]int, 10)
	s2 := make([]int, 10)
	//还可以指定容量
	s3 := make([]int, 10, 15)
	fmt.Println(s, s1, s2, s3)

	//初始化切片
	arr := []int{1, 2, 3, 4, 5}
	//array := [...]int{1,2,3,4,5}
	//由已有的数组初始化(用切片也是可以的)，[:]就表示整个数组，当然，也可以[startIndex:],[:endIndex]
	//arr1 := array[0:3]
	arr1 := arr[1:3]
	//len()获取切片长度,cap()获取切片容量,不指定的话，默认容量是等于长度的，不过如果是arr1这种的话,它的容量是arr的容量
	fmt.Println(arr1, len(arr1), cap(arr1))
	fmt.Println(len(arr), cap(arr))

	//空(nil)切片，切片未初始化时，默认都是0，值是[]，长度是0，容量是0
	fmt.Println(s, len(s), cap(s))

	//截取切片（上面初始化切片里用的方式就是截取切片来初始化的，不多说）

	//切片函数
	//append，追加元素
	//追加空切片
	s = append(s, 0)
	fmt.Println(s)
	s = append(s, 1)
	fmt.Println(s)
	//追加多个元素
	s = append(s, 2, 3, 4)
	fmt.Println(s)

	//扩容
	//创建一个两倍容量的切片
	ns := make([]int, len(s), 2*cap(s))
	//拷贝
	copy(ns, s)
	fmt.Println(ns)
}
