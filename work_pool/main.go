package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker 负责处理任务
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d 开始处理任务 %d\n", id, job)

		time.Sleep(time.Second)

		fmt.Printf("Worker %d 完成任务 %d\n", id, job)
	}
}

func main() {
	const workerCount = 5
	const jobCount = 20

	jobs := make(chan int, jobCount)

	var wg sync.WaitGroup

	// 创建固定数量的 Worker
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go worker(i, jobs, &wg)
	}

	// 投递任务
	for i := 1; i <= jobCount; i++ {
		jobs <- i
	}

	// 通知 Worker：没有更多任务了
	close(jobs)

	// 等待所有 Worker 完成
	wg.Wait()

	fmt.Println("所有任务处理完成")
}
