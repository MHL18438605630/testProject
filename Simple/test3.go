package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	updateNums(&nums)
	fmt.Println(nums)

}
func updateNums(nums *[5]int) {
	nums[1] = 20
}
