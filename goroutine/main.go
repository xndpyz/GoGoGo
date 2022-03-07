package main

import "fmt"

func main() {
	//c := make(chan int, 5)

	//执行顺序 先存5个, 取出一个，再存一个
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//}()
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c)
	//}

	//close 用法 结合for range 的时候使用

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
