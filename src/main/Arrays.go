/**
 * @author  ironion
 * @date  2020/5/10 11:34 上午
 * @note
 */
package main

import "fmt"

func main() {

	//数组定义
	//未初始化
	var array [10]int8
	array[0] = 1
	fmt.Println(array)
	//初始化（指定大小），{}里的数量不能大于指定的大小
	var array2 = [3]int8{1, 2, 3}
	fmt.Println(array2)
	//初始化（不指定大小）
	var array3 = [...]int8{1, 2, 3, 3, 4, 2, 3}
	fmt.Println(array3)

	//顺带说一句 int8 表示int的范围是2的8次方，就是到256，其他的类似
	//下面例子可以看出对数组的拷贝是深拷贝，对复制后的数组的操作不会影响到原有的数组
	a := array3[1]
	b := array
	fmt.Println(a, "...", array3[1])
	fmt.Println(b, "...", array)
	a = 4
	b[0] = 15
	fmt.Println(a, "...", array3[1])
	fmt.Println(b, "...", array)

}
