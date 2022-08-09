package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 用于计算执行操作次数
	var ops int64

	// reads 和 writes 通道分别被其他 Go 协程用来分布读和写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 私有的state，可利用其他线程来修改state
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				a := <-read.resp
				fmt.Println("a = ", a)
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				b := <-write.resp
				fmt.Println("b = ", b)
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	time.Sleep(time.Second)

	opsFinalVal := atomic.LoadInt64(&ops)

	fmt.Println("ops : ", opsFinalVal)
}
