package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	//io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。
	//o.Reader 接口有一个 Read 方法：
	//func (T) Read(b []byte) (n int, err error)
	//Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
	//初始化一个需要被读取的值
	read := strings.NewReader("wqrwqewqrgb  wrqeqw12324,2321")
	//每次以8字节读取
	bytes := make([]byte, 8)
	//用个死循环读取
	for {
		n, err := read.Read(bytes)
		fmt.Println(n, err, bytes)
		fmt.Printf("%q \n", bytes[:n])
		//判断是不是读到末尾了
		if err == io.EOF {
			fmt.Println("over...")
			break
		}
	}
}

//有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。
//例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。
