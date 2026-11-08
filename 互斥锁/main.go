package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var x = int(0)
var lock = sync.Mutex{}

func add(*sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 500000; i++ {
		// 上锁
		lock.Lock()
		x += 1
		// 解锁
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add(&wg)
	go add(&wg)
	wg.Wait()
	fmt.Println(x)
}
