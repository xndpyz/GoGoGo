package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//循环打印cat, dog, fish 按顺序，各100次
	echoAnimals()

	//交替打印英文
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	sign := make(chan struct{})

	go echoNums(10, ch1, ch2)
	go echoEnglish("ABCDEFGHIJ", ch2, ch1, sign)
	ch1 <- struct{}{}
	//time.Sleep(1 * time.Second)
	<-sign
}

func echoNums(num int, in, out chan struct{}) {
	for i := 0; i < num; i += 2 {
		<-in
		fmt.Println(i)
		fmt.Println(i + 1)
		out <- struct{}{}
	}
}

func echoEnglish(str string, in, out, end chan struct{}) {
	for i := 0; i < len(str); i += 2 {
		if i > 5 {
			end <- struct{}{}
		}
		<-in
		fmt.Println(str[i])
		fmt.Println(str[i+1])
		out <- struct{}{}
	}
	//end <- struct{}{}
}

func echoAnimals() {
	var catCount uint64
	var dogCount uint64
	var fishCount uint64
	var wg sync.WaitGroup
	wg.Add(3)
	dogch := make(chan struct{})
	catch := make(chan struct{})
	fishch := make(chan struct{})

	go echoDog(&wg, dogCount, dogch, catch)
	go echoCat(&wg, catCount, catch, fishch)
	go echoFish(&wg, fishCount, fishch, dogch)

	dogch <- struct{}{}
	fmt.Println(dogCount, 1111)
}

func echoDog(wg *sync.WaitGroup, count uint64, dogch, catch chan struct{}) {
	for {
		if count >= 100 {
			wg.Done()
			return
		}
		<-dogch
		fmt.Println("dog")
		atomic.AddUint64(&count, 1)
		catch <- struct{}{}
	}
}

func echoCat(wg *sync.WaitGroup, count uint64, catch, fishch chan struct{}) {
	for {
		if count >= 100 {
			wg.Done()
			return
		}
		<-catch
		fmt.Println("cat")
		atomic.AddUint64(&count, 1)
		fishch <- struct{}{}
	}
}

func echoFish(wg *sync.WaitGroup, count uint64, fishch, dogch chan struct{}) {
	for {
		if count >= 100 {
			wg.Done()
			return
		}
		<-fishch
		fmt.Println("fish")
		atomic.AddUint64(&count, 1)
		dogch <- struct{}{}
	}
}
