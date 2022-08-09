package main

import "fmt"

type Student struct {
	name string
	age  int
	sex  string
}

func main() {

	stu1 := Student{"MHL", 23, "girl"}
	stu := updateStudentInfo(&stu1)
	fmt.Println(stu1)
	fmt.Println(stu)

}

func updateStudentInfo(stu *Student) *Student {
	stu.name = "hhhhhh"
	stu.sex = "boy"
	stu.age = 29
	return stu
}
