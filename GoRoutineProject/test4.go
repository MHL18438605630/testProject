package main

import (
	"fmt"
	"time"
)

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func chanTest() {
	ch := pump()

	go func() {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				fmt.Println("break ...")
				break
			}
		}
	}()

	select {
	case <-time.After(time.Second):
		return
	}

}

func main() {
	chanTest()
}
