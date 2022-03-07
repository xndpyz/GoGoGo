package main

import (
	"fmt"
	"sync"
)

func main() {
	//sync
	//读写锁(读多写少的情况	), 互斥锁

	//只执行一次, 用了互斥锁， sync.Mutex
	o := &sync.Once{}

	for i := 0; i < 5; i++ {
		o.Do(func() {
			fmt.Println(i)
		})

	}

}
