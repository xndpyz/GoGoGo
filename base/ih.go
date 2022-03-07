package main

import "fmt"

func main() {
	//* 指明地址, &取出地址
	var a int
	a = 2
	var b *int
	b = &a
	*b = 999
	fmt.Println(a, b)
	fmt.Println(a == *b, &a == b) // true true

	//可以修改数组的值
	//阻塞等待
}
