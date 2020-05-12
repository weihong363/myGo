package main

import (
	"fmt"
	"time"
)

func main() {
	//并发，Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
	//goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
	//写法为：go 函数名( 参数列表 )，例如：go func(1,2,3)
	//Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
	//同一个程序中的所有 goroutine 共享同一个地址空间。
	go run("A")
	run("B")

	//通道channel，是用来传递数据的一个数据结构。
	//通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
	//声明通道
	c := make(chan int)
	//默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
	s := []int{2, 4, 3, 7, 3, 6, -9, -3}
	go sum(s[:4], c)
	go sum(s[4:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)

	//通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：c := make(chan int,1000)
	//带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
	//不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
	//注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。
	//接收方在有值可以接收之前会一直阻塞。
	//c1 := make(chan int,1)
	//c1 <- 1
	//c1 <- 2
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//因为这边只设了缓冲大小为1，这时候两个线程去争用资源，就导致了死锁
	//下面改成2就可以了
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	//通道不能一直开着，当不需要的时候就要把它关闭掉
	c3 := make(chan int, 10)
	go fibonacci(cap(c3), c3)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c3 {
		fmt.Printf("%d ", i)
	}
}

func run(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("thread" + s + "is running...")
	}
}

/**
 * @description 计算切片所有数值之和，然后传到通道
 * @date 2020/5/13 12:54 上午
 * @author ironion
 * @param
 * @return
 **/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

/**
 * @description 带通道关闭的斐波那契数列
 * @date 2020/5/13 1:05 上午
 * @author ironion
 * @param
 * @return
 **/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
