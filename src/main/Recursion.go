package main

import "fmt"

func main() {
	//递归recursion
	//斐波那契数列
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
