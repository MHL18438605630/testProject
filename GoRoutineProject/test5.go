package main

import (
	"fmt"
	"time"
)

func chanTest1() {
	tickCount := time.Tick(time.Second)
	boom := time.After(5 * time.Second)

	count := 0
	for {
		select {
		case <-tickCount:
			count++
			fmt.Println(count)
		case <-boom:
			fmt.Println("BOOM!!!")
			return
		}
	}

}

func main() {

	chanTest1()

}
