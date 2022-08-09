package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A的i = ", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B的i = ", i)
	}
}

func main() {

	runtime.GOMAXPROCS(8)
	go a()
	go b()

	time.Sleep(time.Second)

}
