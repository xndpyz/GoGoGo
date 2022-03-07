package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	g  int
	mu sync.Mutex
	//rmu sync.RWMutex
)

func main() {
	// go 读写锁和并发锁

	//demo1()

	demo2()
}

func demo1() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("func1:....", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("func2:....", i)
		}
	}()
	wg.Wait()

}

func demo2() {
	wg.Add(2)
	go func() {
		for i := 0; i < 50000; i++ {
			mu.Lock()
			g += 1
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 50000; i++ {
			mu.Lock()
			g += 1
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(g)
}

func demo3() {
	// 同步锁和读写锁区别

}

type single struct {
}

var instance *single

// 并不是并发安全的
func singleton() *single {
	if instance == nil {
		return &single{}
	}
	return instance
}

func singletonNew() *single {
	var once sync.Once
	once.Do(func() {
		instance = &single{}
	})
	return instance
}
