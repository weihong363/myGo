package main

import "fmt"

func main() {
	//range，其实就是求个范围，和for结合使用
	s := []int{1, 2, 3}
	for i2, i3 := range s {
		//看下这里第一个和第二个变量是啥子（其实猜也知道，一个是索引，一个是对应的值）
		fmt.Println(i2, i3)
	}
	sum := 0
	//当我们不需要索引的时候可以这样，_为空白符
	for _, i3 := range s {
		sum += i3
	}
	fmt.Println(sum)
	//这时可以想，一个索引，一个值，那这岂不是可以用在遍历map上，一个key，一个value岂不美哉
	m := map[string]string{"a": "zhangsan", "b": "lisi"}
	for i2, i3 := range m {
		fmt.Println(i2, i3)
	}
	//除了切片（数组），map之外，还可以用来枚举unicode字符串
	for i2, i3 := range "hellow" {
		fmt.Println(i2, i3)
	}
}
