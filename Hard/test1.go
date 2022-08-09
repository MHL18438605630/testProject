package main

import (
	"fmt"
	"runtime"
)

func ReturnExit() {
	func() {
		fmt.Println("开心就好")
		runtime.Goexit()
	}()

	fmt.Println("kkkkkkkkk")

}

func main() {

	ReturnExit()
}
