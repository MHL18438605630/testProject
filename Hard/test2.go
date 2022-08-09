package main

import (
	"fmt"
	"time"
)

//func NewTask() {
//	i := 0
//	for {
//		i++
//		fmt.Println("子协程中 i = ", i)
//		time.Sleep(1 * time.Second)
//	}
//
//}

func main() {
	//go NewTask()
	//i := 0
	//for {
	//	i++
	//	fmt.Println("主协程中 i = ", i)
	//	time.Sleep(1 * time.Second)
	//
	//}

	go func(a, b int) {
		for {
			fmt.Println("a =", a, ",b = ", b)
			a++
			b++
			time.Sleep(1 * time.Second)
		}
	}(10, 20)

	i := 0
	for {
		i++
		fmt.Println("主协程中 i = ", i)
		time.Sleep(1 * time.Second)
	}

}
