package main

import "fmt"

func SubGoroutine(i int, done chan bool) {
	fmt.Printf("子协程%d\n", i)
	done <- true
}

func main() {
	fmt.Println("main协程开始")
	done := make(chan bool)
	go SubGoroutine(1, done)
	<-done
	fmt.Println("main协程结束")
}
