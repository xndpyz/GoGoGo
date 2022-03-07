package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	1. sync
*/

func main() {
	day6()
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
