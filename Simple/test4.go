package main

import "fmt"

func main() {
	var sliNums1 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d sliNums1=%v\n", len(sliNums1), cap(sliNums1), sliNums1)

	sliNums1 = append(sliNums1, 6)
	fmt.Printf("len=%d cap=%d sliNums1=%v\n", len(sliNums1), cap(sliNums1), sliNums1)

	sliNums1 = append(sliNums1, 7)
	fmt.Printf("len=%d cap=%d sliNums1=%v\n", len(sliNums1), cap(sliNums1), sliNums1)

	sliNums1 = append(sliNums1[:3], sliNums1[5:]...)
	fmt.Printf("len=%d cap=%d sliNums1=%v\n", len(sliNums1), cap(sliNums1), sliNums1)

}
