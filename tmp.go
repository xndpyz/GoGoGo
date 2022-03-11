package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	1. sync
*/

func main() {
	// go 强类型预约 uint类型
	//int 在32操作系统是4个字节，64位是8个字节
	var a uint = 1
	var b uint = 2

	// 如果是32系统 负数的计算都是补码，2的32次方-1
	fmt.Println(a - b)
	//fmt.Println(math.Pow(2, 64))

	/*
		rune 类型 int32类似，可以计算特定符号的长度，比如中文

	*/
	var num uint64 = 1
	atomic.AddUint64(&num, 1)

	//day6()
	//day7()

	//echoPrime()

	ch := make(chan int)

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- 10000

}

func echoPrime() {
	ch := make(chan int)
	go generate(50, ch)
	for {
		prime := <-ch
		fmt.Printf("%d->", prime)
		ch2 := make(chan int)
		go filter(prime, ch, ch2)
		ch = ch2
	}

}

func filter(prime int, in, out chan int) {
	//fmt.Println("process")
	for {
		num := <-in
		if num%prime != 0 {
			out <- num
		}
	}
}

func generate(num int, ch chan int) {
	for i := 2; i < num+1; i++ {
		ch <- i
	}
	close(ch)
}

func day7() {
	// 交替打印英文字母
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	signal := make(chan struct{})
	go english("ABCDEFGHIJKL", ch1, ch2)
	go nums(10, ch2, ch1, signal)
	ch1 <- struct{}{}
	<-signal
}

func english(str string, in, out chan struct{}) {
	for i := 0; i < len(str); i += 2 {
		<-in
		fmt.Println(str[i])
		fmt.Println(str[i+1])
		out <- struct{}{}
	}
}

func nums(num int, in, out, sign chan struct{}) {
	for i := 1; i < num; i += 2 {
		<-in
		fmt.Println(i)
		fmt.Println(i + 1)
		out <- struct{}{}
		if i == num-1 {
			sign <- struct{}{}
		}
	}
}

func day6() {
	// sync.Mutex, sync.RMutex, sync.Once()
	// sync.RMutex 读多写少的情况

	// Once 底层结构体由Mutex互斥锁组成，只执行一次
	var o sync.Once
	for i := 0; i < 5; i++ {
		o.Do(func() {
			fmt.Println("sync.Once 执行次数", i)
		})
	}

	//l := &sync.RWMutex{}
	//go echo(l)
	//go echo(l)
	//go echo(l)
	//for {
	////} // 程序阻塞在这

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("waitGroup 1")
	}()
	wg.Wait()

	noPeatStr("abbcd")
}

func echo(l *sync.RWMutex) {
	l.Lock()
	fmt.Println("echo： start")
	time.Sleep(1 * time.Second)
	l.Unlock()
}

func noPeatStr(s string) int {
	right, res := 0, 0
	hash := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(hash, s[i])
		}
		for right < len(s) && hash[s[right]] == 0 {
			hash[s[right]]++
			right++
		}
		res = max(right-i-1, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
