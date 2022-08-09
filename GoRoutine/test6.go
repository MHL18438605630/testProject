package main

import "fmt"

func main() {
	done := make(chan bool) // 标记量
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i = ", i)
		}
		done <- true // 利用通信进行共享内存
	}()
	<-done
}
