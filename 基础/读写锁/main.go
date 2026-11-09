package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg = sync.WaitGroup{}

var rwLock sync.RWMutex

func read() {
	rwLock.RLock()
	time.Sleep(1 * time.Millisecond)
	rwLock.RUnlock()
	wg.Done()
}

func write() {
	rwLock.Lock()
	x += 1
	time.Sleep(10 * time.Millisecond)
	rwLock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}

	fmt.Println(time.Now().Sub(start))
}
