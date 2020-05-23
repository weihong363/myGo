package main

import (
	"io"
	"os"
	"strings"
)

//go指南上的习题
//练习：rot13Reader
//有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。
//例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。
//编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。
//rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

//要求自己实现的部分（参考自某博客的代码）

//rot13算法是把当前字母替换成它往后移13位的加密算法
func rot(out byte) byte {
	switch {
	case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
		out += 13
	case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
		out -= 13
	}
	return out
}

//重写Read方法，遍历输入的byte，对每个元素进行替换
func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot(b[i])
	}
	return n, err
}
