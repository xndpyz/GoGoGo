package main

import "fmt"

func main() {
	fmt.Println("We are all consenting adults here")
	fmt.Println("利用协程打印素数demo")

	ch := make(chan int)
	go generateNum(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filterNum(ch, ch1, prime)
		ch = ch1
	}
}

func generateNum(ch chan int) {
	for i := 2; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func filterNum(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
