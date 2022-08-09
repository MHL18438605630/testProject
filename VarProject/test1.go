package main

import "fmt"

func main() {

	//var b [3]int = [3]int{1,2,3}

	//var c [3]string =[3]string{"1","2","3"}

	var e map[int]string = map[int]string{1: "a", 2: "b", 3: "c"}

	fmt.Println(" e = ", e[1])

}
