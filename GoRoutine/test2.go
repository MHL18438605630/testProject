package main

import (
	"fmt"
	"time"
)

func ReadChanVal(chanVal chan int) {
	for {
		a := <-chanVal
		fmt.Println("a = ", a)
	}
}

func WriteChanVal(chanVal chan int, intVal int) {
	chanVal <- intVal
}

func main() {

	mainChanVal := make(chan int)

	go WriteChanVal(mainChanVal, 1)

	go ReadChanVal(mainChanVal)

	time.Sleep(time.Second)

}
