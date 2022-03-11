package main

import (
	"fmt"
	"sync"
)

func main() {
	//tmp1()
	//fmt.Println('a' - 'b')

	//tmp2()

	tmp3("abbcdfifgjj")

	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	fmt.Println(a, x, y)
	x = append(x, y...) // [0], [2, 3]
	fmt.Println(a, x, y)
	x = append(x, y...)
	fmt.Println(a, x)
}

func tmp3(s string) int {

	left, right := 0, 1
	max := 0
	hash := map[byte]bool{}
	hash[s[0]] = true
	for right < len(s) {
		for hash[s[right]] {
			delete(hash, s[left])
			left++
		}
		if right-left+1 > max {
			max = right - left + 1
		}
		hash[s[right]] = true
		right++
	}
	fmt.Println(max)
	return 0
}

func tmp2() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	sign := make(chan struct{}, 1)

	go echo1(ch1, ch2)
	go echo2(ch2, ch1, sign)
	ch1 <- struct{}{}
	<-sign
}

func echo1(in, out chan struct{}) {
	for i := 0; i < 26; i += 2 {
		<-in
		fmt.Println(i + 1)
		fmt.Println(i + 2)
		out <- struct{}{}
	}
}

func echo2(in, out, sign chan struct{}) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(str); i += 2 {
		<-in
		fmt.Println(string(str[i]))
		fmt.Println(string(str[i+1]))
		out <- struct{}{}
		if i == 24 {
			sign <- struct{}{}
		}
	}

}

func tmp1() {
	wg := sync.WaitGroup{}
	//wg.Add(3)
	//
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)
	sighal := make(chan struct{})

	go cat(ch1, ch2, &wg)
	go fish(ch2, ch3, &wg)
	go dog(ch3, ch1, sighal, &wg)

	ch1 <- struct{}{}
	<-sighal
	wg.Wait()
}

func cat(in, out chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Println("cat", i)
		out <- struct{}{}
		//if i == 99 {
		//	wg.Done()
		//}
	}
}

func fish(in, out chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Println("fish", i)
		out <- struct{}{}
		//if i == 99 {
		//	wg.Done()
		//}
	}
}

func dog(in, out, signal chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Println("dog", i)
		out <- struct{}{}
		if i == 99 {
			signal <- struct{}{}
		}
	}
}
