package main

import (
	"fmt"
	"sort"
)

type UserInfo struct {
	userName string
	userAge  int
}

type userInfoNums []UserInfo

func (u userInfoNums) Len() int {
	return len(u)
}

func (u userInfoNums) Less(i, j int) bool {
	return u[i].userAge > u[j].userAge
}

func (u userInfoNums) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {

	users := userInfoNums{{"MHL", 25}, {"ABC", 24}, {"DEF", 34}}
	sort.Sort(users)
	fmt.Println("users = ", users)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Sort(sort.IntSlice(nums))

	fmt.Println("nums = ", nums)

}
