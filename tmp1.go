package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 理解有缓冲和无缓冲的区别
	// 互斥锁和读写锁

	//t2()

	//t3()

	t5()
}

func t2() {
	ch := make(chan int)
	//mu := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func() {
			//mu.Lock()
			fmt.Println("获取：", <-ch)
			//mu.Unlock()
			time.Sleep(1 * time.Second)
		}()
	}
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	time.Sleep(3 * time.Second)
}

func t1() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < cap(ch); i++ {
			ch <- i
		}
	}()

	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}

func t3() {
	ch := make(chan int, 1)
	ch <- 1 // 无缓冲通道：发送必须要有接受, 否则会阻塞接受
	go func() {
		fmt.Println(<-ch)
	}()
}

func t4() {
	var lock sync.RWMutex
	lockFunc(&lock)
	lockFunc(&lock)
	lockFunc(&lock)
	for {
	} // 程序一直阻塞在这
}

func t5() {
	flag := make(chan bool)
	message := make(chan int)

	go child(flag, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	flag <- true
	fmt.Println("主线程结束")
}

func child(flag chan bool, mes chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-mes:
			fmt.Println("接收到值：", m)
		case <-flag:
			fmt.Println("子线程结束")
			//return
		}
	}
}

func lockFunc(lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("读写锁")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

// sync.Once()
func lockFuncTwo(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("互斥锁")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}
