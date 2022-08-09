package main

import (
	"fmt"
	"time"
)

func main() {

	myTicker := time.NewTicker(time.Second * 3)

	for {
		t1 := <-myTicker.C
		fmt.Println("3秒后，开始执行程序,", t1)
	}
	myTicker.Stop()

}
