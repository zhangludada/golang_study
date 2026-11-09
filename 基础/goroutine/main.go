package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	fmt.Println("hello world")
	// wg计数器减1
	defer wg.Done()
}

// 申明一个wg的全局变量
var wg sync.WaitGroup

func main() {
	// 只占用一个cpu
	runtime.GOMAXPROCS(1)
	// wg的计数器+1
	wg.Add(1)
	go hello()
	fmt.Println("hello main")
	// 等待wg计数器归零
	wg.Wait()
}
