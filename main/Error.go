package main

import "fmt"

func main() {
	//错误处理
	//error是一个接口，定义如下
	//type error interface {
	//    Error() string
	//}
	fmt.Println(throwE(1))
	fmt.Println(throwE(-1))
}

/**
 * @description 重写异常方法
 * @date 2020/5/13 12:39 上午
 * @author ironion
 * @param
 * @return
 **/
func Error() string {
	return "异常！"
}

/**
 * @description 如果输入值小于0抛出error
 * @date 2020/5/13 12:32 上午
 * @author ironion
 * @param
 * @return
 **/
func throwE(value int) (int, string) {
	if value < 0 {
		return value, Error()
	}
	return value, ""
}
