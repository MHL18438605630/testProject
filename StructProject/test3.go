package main

import "fmt"

type info struct {
	addr string
	name string
}

type info2 struct {
	addr string
	name string
}

type User1 struct {
	id int32
	info
	iii info2
}

func (pInfo *info) printInfo() {
	fmt.Println(pInfo)
}

func (pInfo *info2) printInfo2() {
	fmt.Println(pInfo)
}
func (pUser *User1) printUser() {
	fmt.Println(pUser)
}

func main() {
	user := User1{
		1,
		info{
			"aaa",
			"baba",
		},
		info2{
			"ccc",
			"cbc",
		},
	}

	user.printInfo()
	user.printUser()
	user.iii.printInfo2()

}
