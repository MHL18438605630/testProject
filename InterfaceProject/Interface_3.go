package main

import (
	"fmt"
	"strconv"
)

type UserInter interface {
	GetUserName() string
	SetUserName(newName string)
	PrintUser() string
}

type UserInf struct {
	userName string
	userAge  int
	userSex  string
}

func (u UserInf) GetUserName() string {
	//TODO implement me
	return u.userName
}

func (u *UserInf) SetUserName(newName string) {
	//TODO implement me
	u.userName = newName
}

func (u UserInf) PrintUser() string {
	//TODO implement me
	useInfo := u.userName + ", " + strconv.Itoa(u.userAge) + ", " + u.userSex
	return useInfo
}

func main() {
	user1 := UserInf{
		"MHL",
		25,
		"Woman",
	}

	user2 := &UserInf{
		"DEF",
		23,
		"Man",
	}

	fmt.Println(user1.PrintUser())
	user1.SetUserName("ABC")

	fmt.Printf(user1.PrintUser())

	fmt.Println()

	fmt.Println(user2.PrintUser())
	user2.SetUserName("GHI")

	fmt.Printf(user2.PrintUser())
}
