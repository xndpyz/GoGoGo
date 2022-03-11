package main

import "fmt"

func main() {
	// 考察切片和数组的区别
	// 切片的扩容机制
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a) //100, 200, 3
		}
		a[k] = 100 + v
	}
	fmt.Println(a) //101, 102, 103 //101 300 103

	b := a
	b[0] = 1000
	fmt.Println(a, b)

}
