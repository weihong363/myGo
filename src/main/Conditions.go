package main

import "fmt"

func main() {
	//if..else和java没啥不同，少了括号而已
	if true {

	} else if true {

	} else {

	}

	expr := 1
	//switch也差不多其实
	switch expr {
	case 1:
		fmt.Println(1)
		//break
		//fallthrough
	case 2:
		fmt.Println(2)
	case 3:
	default:
	}
	//这里要注意下
	//switch支持多条件匹配，像这样 case 1,2,3:
	//每个case默认隐式添加了break，只会执行匹配到的case
	//如果想执行多个case，需要加fallthrough关键字
	//如：
	//case1:
	//fallthrough
	//case2
	//这样case1执行完后还会去走case2的操作，除非你在case1里显式地加break

	//select语句，类似于switch。select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。
	//一个默认的子句应该总是可运行的
	//每个 case 都必须是一个通信
	//所有 channel 表达式都会被求值
	//所有被发送的表达式都会被求值
	//如果任意某个通信可以进行，它就执行，其他被忽略。
	//如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
	//否则：
	//如果有 default 子句，则执行该语句。
	//如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

	var x, y int
	//channel 是 goroutine 之间通信的一种方式，可以类比成 Unix 中的进程的通信方式管道。
	//可以写成这样a := make(chan int)
	var a, b, c chan int
	select {
	case x = <-a:
		fmt.Println("recieve ", x, " from a\n")
	case b <- y:
		fmt.Println("send ", y, " to b\n")
	case z, ok := (<-c):
		if ok {
			fmt.Println("recieve ", z, " from c\n")
		} else {
			fmt.Println("c is closed\n")

		}
	default:
		fmt.Println("no communication")
	}
}
