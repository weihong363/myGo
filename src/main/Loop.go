package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	collection := []string{"testmodule", "bbb", "ccc"}
	for i2, i3 := range collection {
		//i2标识数组中的下标，i3标识数组中的值
		fmt.Println(i2, i3)
	}
	//死循环就
	//for{
	//	fmt.Println("")
	//}

	//循环控制语句continue和break就不多说了，这里还多了个goto，如：
	a := 10
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 2
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}

}
