package main

import (
	"fmt"
	"time"
)

func main() {
	ChanVal1 := make(chan int)
	ChanVal2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		<-ChanVal2    // ===> 发送数据
		ChanVal1 <- 1 // ===> 接收数据
	}()

	select {
	case v := <-ChanVal1:
		fmt.Printf("%d <- ChanVal1", v)
	case ChanVal2 <- 1:
		fmt.Println("ChanVal2 <- 1")
	}

}
