package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//runtime.GOMAXPROCS(1)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for true {
	//		//fmt.Println("hello  子协程！！！")
	//	}
	//}()
	//wg.Wait()
	//fmt.Println("主协程  ")

	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("hello go!")

}
