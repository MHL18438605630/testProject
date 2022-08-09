package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Goroutine is done")
	}()

	timer.Stop()
	timer.Reset(5 * time.Second)

	select {
	case <-time.After(10 * time.Second):

	}

}
