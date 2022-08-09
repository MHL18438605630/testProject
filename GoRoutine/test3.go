package main

import (
	"fmt"
	"time"
)

func SendChanVal(chanVal chan int) {
	fmt.Println("ChanVal = ", <-chanVal)
}

func main() {

	chanVal := make(chan int)

	//

	go SendChanVal(chanVal)
	chanVal <- 123 // 接收者
	time.Sleep(time.Second)
}
