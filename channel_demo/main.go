package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// ch2被赋值
func f1(ch2 chan<- int) {
	for i := 0; i < 10; i++ {
		ch2 <- i
	}
	close(ch2)
}

// ch2取出值，ch3被赋值
func f2(ch2 <-chan int, ch3 chan<- int) {
	// 方法一
	//for {
	//	tmp, ok := <-ch2
	//	if !ok {
	//		break
	//	}
	// 方法二 ， 建议用这种，不用谢错误判断
	for tmp := range ch2 {
		ch3 <- tmp * tmp
	}
	close(ch3)
}

func main() {
	// 无缓冲区通道
	ch := make(chan int, 1)
	ch <- 10
	x := <-ch
	wg.Add(1)
	go func() {
		fmt.Println(x)
		wg.Done()
		defer close(ch)
	}()
	wg.Wait()

	// 带缓冲区通道
	ch1 := make(chan int, 1)
	ch1 <- 20
	y := <-ch1
	fmt.Println(y)

	// 练习，把ch1的平方，存入ch2
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	// 给ch1传递值
	go f1(ch2)
	// 给ch2传递值
	go f2(ch2, ch3)

}
