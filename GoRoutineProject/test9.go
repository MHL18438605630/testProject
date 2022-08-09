package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("&quot;worker:%d start job:%d\n&quot", id, j)
		time.Sleep(time.Second)
		fmt.Printf("&quot;worker:%d end job:%d\n&quot;", id, j)
		results <- j * 2
	}

}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		fmt.Println("Result = ", <-results)
	}

}
