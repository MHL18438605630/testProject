package main

import "fmt"

type User struct {
	userName string
	userAge  int
}

func (u *User) PrintUser() {
	fmt.Println(u.userName, u.userAge)
}
func main() {

	user := User{
		"MHL",
		25,
	}

	user.PrintUser()

}
