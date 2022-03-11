package main

import (
	"fmt"
	"sync"
)

func main() {
	//交叉打印cat, dog, fish 各100次
	f1()

	//交叉打印26位英文字母
}

func f1() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	go cat(&wg, ch1, ch2)
	go dog(&wg, ch2, ch3)
	go fish(&wg, ch3, ch1)

	ch1 <- struct{}{}

	wg.Wait()
}

func cat(wg *sync.WaitGroup, in, out chan struct{}) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Println("cat", i)
		out <- struct{}{}
		if i == 99 {
			wg.Done()
			return
		}
	}
}

func dog(wg *sync.WaitGroup, in, out chan struct{}) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Println("dog", i)
		out <- struct{}{}
		if i == 99 {
			wg.Done()
			return
		}
	}
}

func fish(wg *sync.WaitGroup, in, out chan struct{}) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Println("fish", i)
		out <- struct{}{}
		if i == 99 {
			wg.Done()
			return
		}
	}
}
