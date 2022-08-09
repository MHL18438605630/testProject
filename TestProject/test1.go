package main

import "fmt"

var pc [10]int

func init() {
	for i := 0; i < 10; i++ {
		pc[i] = i
	}
}

func main() {

	fmt.Println("pc[1] + pc[2] = ", pc[1]+pc[2])

}
