package main

import "fmt"

var ChanVal chan int = make(chan int)

func NewTask(a, b int) {
	ChanVal <- a + b
}

func main() {
	go NewTask(1, 2)

	date, err := <-ChanVal

	fmt.Println("data =", date, ",err = ", err)

}
