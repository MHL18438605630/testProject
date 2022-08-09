package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string `info:"姓名信息"`
	Age  int    `info:"年龄信息"`
}

func GetTagInfo(input interface{}) {
	InputType := reflect.TypeOf(input).Elem()

	for i := 0; i < InputType.NumField(); i++ {
		TagInfo := InputType.Field(i).Tag.Get("info")
		fmt.Println("Info = ", TagInfo)
	}

}

func main() {
	stu := Stu{"MHL", 24}
	GetTagInfo(&stu)

}
