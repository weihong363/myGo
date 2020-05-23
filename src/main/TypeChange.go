package main

import (
	"fmt"
	"strconv"
)

func main() {
	//类型转换
	//type(expression)
	a := "4"
	b := "5.2"
	c := "0.2"
	//字符串转其他类型要用strconv
	i2, error := strconv.Atoi(a)
	f, error1 := strconv.ParseFloat(b, 32)
	f2, error2 := strconv.ParseFloat(c, 32)
	fmt.Println(i2, error)
	fmt.Println(f, error1)
	fmt.Println(f2, error2)
	//这里i2和f2类型不同，需要转型才能处理
	i3 := float64(i2) / f2
	//数字的转型
	fmt.Println(i3)
}
