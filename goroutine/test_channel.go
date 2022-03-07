package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		defer fmt.Println("传输数据结束")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("正在发送数据：", i)
		}
		close(ch)
	}()
	time.Sleep(1)
	for i := 0; i < 4; i++ {
		data, ok := <-ch
		if ok {
			fmt.Println(data)
		}
	}
}
