package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	go func() {
		// time.Sleep(time.Second)
		c1 <- 1 // 接收数据
	}()
	select {
	case c1 <- 1: // 发送数据
		fmt.Printf("c1 <- 1")
	default:
		fmt.Println("default")
	}
}
