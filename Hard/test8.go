package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // 写入成功
			x = y
			y = x + y
		case <-quit: // 读取成功
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 尝试从c中读取数据，没有数据，则会c阻塞，循环6次
		}
		quit <- 0 // 发送数据0到quit中，表示退出
	}()

	fibonacii(c, quit)
}
