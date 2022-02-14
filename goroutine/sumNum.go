package main

import (
	"fmt"
)

func main() {
	fmt.Println("协程求和")

	goroutine()
}

func sum(start, end int, ch chan int) {
	var res int
	for i := start; i < end; i++ {
		res += i
	}
	ch <- res
}

func goroutine() {
	var res int
	nums := 1000000
	ch := make(chan int, 4)
	num := nums / cap(ch)
	for i := 0; i < cap(ch); i++ {
		go sum(i*num, i*num+num, ch)
	}
	for i := 0; i < cap(ch); i++ {
		//tmp := <- ch
		res += <-ch
	}
	fmt.Println(res + nums)
}
