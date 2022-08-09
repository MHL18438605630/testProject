package main

import "fmt"

type User struct {
	UserName string
	UserAge  int
	UserSex  string
}

type UserInterface interface {
	GetUserName() string
	SetUserName(userName string)
}

func (this *User) GetUserName() string {
	return this.UserName
}

func (this *User) SetUserName(userName string) {
	this.UserName = userName
}

func PrintUser(userInterface UserInterface, usr string) {
	userInterface.SetUserName(usr)
	fmt.Println(userInterface.GetUserName())
}

func main() {

	user1 := &User{
		UserName: "MHL",
		UserAge:  25,
		UserSex:  "Woman",
	}
	//
	//user2 := User{
	//	UserName: "MHL",
	//	UserAge:  25,
	//	UserSex:  "Woman",
	//}
	//
	//fmt.Println(user1 == &user2)

	PrintUser(user1, "ABC")

}
