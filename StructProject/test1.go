package main

import "fmt"

type User struct {
	UserName string
	UserAge  int
}

func (u *User) PrintUser() {
	fmt.Println("User = { ", u.UserName, ",", u.UserAge, " }")
}

func main() {

	user := User{
		UserName: "MHL",
		UserAge:  25,
	}

	user.PrintUser()

}
