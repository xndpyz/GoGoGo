package main

import "fmt"

func main() {
	fmt.Println("通道的多路复用")

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println("存数据....", i)
		case res := <-ch:
			fmt.Println("取出数据....", res)
		default:
			fmt.Println("通道异常")
		}

	}
}
