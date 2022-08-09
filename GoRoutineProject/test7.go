package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var numMap map[int]string
var mapMutex sync.RWMutex
var mapOnce sync.Once

func initMap() {
	numMap = map[int]string{
		1: "MHL",
		2: "ABC",
		3: "DEF",
		4: "GHI",
	}
}

func getMap() {
	mapOnce.Do(initMap)
	mapMutex.RLock()
	for i, s := range numMap {
		fmt.Println("key = ", i, ", value = ", s)
	}
	mapMutex.RUnlock()
}

func setMap(index int, newValue string) {
	mapOnce.Do(initMap)
	_, ok := numMap[index]
	if !ok {
		fmt.Println("not exist")
		return
	}

	mapMutex.Lock()
	numMap[index] = newValue
	mapMutex.Unlock()

}

func main() {

	for i := 1; i < 5; i++ {
		//go getMap()
		go setMap(i, strconv.Itoa(i))
	}

	time.Sleep(1 * time.Second)

}
