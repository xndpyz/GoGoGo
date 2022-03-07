package main

import "fmt"

func main() {
	//byte 是uint8别名，区分字节值和8位无符号整数值
	var ch1 byte = 65
	var ch2 byte = '\x41'
	var ch3 byte = 'A'
	fmt.Println(ch1, ch2, ch3)

	//数组和切片的区别，在编译的时候确定长度
	//深拷贝和浅拷贝，深拷贝copy()，浅拷贝[:]
	//range 便利切片的时候会先保留一个备份

	c := []chips{
		{1},
		{2},
		{3},
		{4},
	}
	for _, v := range c {
		if v.id <= 2 {
			v.id = 3
		}
		fmt.Println(v)
	}
	fmt.Println(c)
}

type chips struct {
	id int
}
