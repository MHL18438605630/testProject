package main

import (
	"fmt"
	"time"
)

// 多协程间通信channel

func readChan(ch chan int) {
	fmt.Println(<-ch)
}

func writeChan(i int, ch chan int) {
	ch <- i
}

func main() {

	ch := make(chan int)
	go writeChan(1, ch)
	go readChan(ch)

	time.Sleep(1 * time.Second)

}
