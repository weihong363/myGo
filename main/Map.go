package main

import "fmt"

func main() {
	//集合Map，这个也很熟了
	//定义的方式有两种，像切片那样，可以通过它的关键字也可以通过make函数来
	var m1 map[string]string
	//m1没法直接赋值，得make之后才行
	m2 := make(map[string]string)
	fmt.Println(m1, m2)
	m2["a"] = "aaaa"
	m2["b"] = "bbbb"
	m2["c"] = "cccc"
	for v := range m2 {
		fmt.Println(m2[v])
	}

	//delete函数，删除指定的键值对
	delete(m2, "c")
	fmt.Println(m2)
}
