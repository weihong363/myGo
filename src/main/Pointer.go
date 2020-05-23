/**
 * @author  ironion
 * @date  2020/5/10 11:21 上午
 * @note
 */
package main

import "fmt"

func main() {
	a := 10
	var b *int
	b = &a
	//输出a的内存地址,和b指针
	fmt.Println(&a)
	fmt.Println(b)
	//使用指针访问变量的值
	fmt.Println(*b)

	//空指针,nil
	var ptr *int
	fmt.Println(ptr)
	//空指针格式化输出的值为0,不能通过*的方式来取，因为是空的
	fmt.Printf("%x\n", ptr)

	//指针数组
	//「因为go中数组是值类型，所以数组传递到其他函数时，其他函数的操作不会影响到原有数组的值
	//比如：数组A调函数把里面的值都加1，在函数里是都加1了。但是传出来的值还是原有的」
	//「」内为某文章里的说法，以下例子实测后发现会改变。。。
	fmt.Println("==================")
	const MAX = 3
	//初始化一个数组
	array := []int{1, 2, 3}
	fmt.Println("传值前")
	fmt.Println(array)
	add(array)
	fmt.Println("传值后")
	fmt.Println(array)
	for i := 0; i < MAX; i++ {
		fmt.Println(array[i])
	}

	//这是某文章里的做法
	arr3 := [3]int{1, 2, 3}
	fmt.Println("测试前", arr3)
	testValue(arr3)
	fmt.Println("测试后", arr3)

	//对比了下，发现主要是处理array的类型不同，[]int和[3]int。这是因为[3]int是数组类型，[]int是切片类型。
	//切片类型调函数时传递了引用，而数组则传了值
	//Go语言中，在函数调用时，引用类型（slice、map、interface、channel）都默认使用引用传递

	//说回正题，指针数组据W3C上说，是为了保存原有的数组用的，具体操作如下
	//初始化一个空指针数组
	var ptrs [MAX]*int
	for i := 0; i < MAX; i++ {
		ptrs[i] = &arr3[i]
	}
	fmt.Println(*ptrs[0], *ptrs[1], *ptrs[2])

	//=========================================

	//指向指针的指针（类似C的二级指针），即指针变量存的是另一个指针变量的地址
	//作用主要体现在函数调用这块，用来改变被调函数中主调函数的指针指向
	var pptr **int
	pptr = &ptrs[0]
	//二级指针地址，一级指针地址，值
	fmt.Println(pptr, *pptr, **pptr)

	//==========================================
	//作为函数参数，例如交换变量的值
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(i2 *int, i22 *int) {
	*i2, *i22 = *i22, *i2
}

func testValue(araya [3]int) {
	//函数内部，修改第一个元素值
	araya[0] = 100
	fmt.Println("测试中", araya)
}

func add(arr []int) []int {
	//函数内部，修改第一个元素值
	for i := 0; i < len(arr); i++ {
		arr[i]++
	}
	fmt.Println("函数中", arr)
	return arr
}
