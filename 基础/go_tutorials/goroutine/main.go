package main

import (
	"fmt"
	"sync"
)

// 等待组
var wg = sync.WaitGroup{}
// 读写锁 
var m = sync.RWMutex{}

func sayHello(index int) {
	// wg-1
	defer wg.Done()
	m.RLock()
	fmt.Printf("hello %d\n", index)
	m.RUnlock()
}

func main() {
	for i := 1; i < 1000; i++ {
		// wg+1
		wg.Add(1)
		// 一般来说wg.Add要写在wg.Done()外面
		go sayHello(i)
	}
	wg.Wait()
	fmt.Println("over")
}

