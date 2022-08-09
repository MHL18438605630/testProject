package main

import (
	"fmt"
	"reflect"
)

type Student1 struct {
	Name string
	Sex  string
	Age  int
}

func (this *Student1) Show() {
	fmt.Println("----打印----")
	fmt.Printf("%s = %v\n", "Student", this)
}

func GetTypeValue(input interface{}) {
	InputTypes := reflect.TypeOf(input)
	fmt.Printf("type=%v\n", InputTypes)
	InputValues := reflect.ValueOf(input)
	fmt.Printf("value=%v\n", InputValues)

	for i := 0; i < InputTypes.NumField(); i++ {
		Field := InputTypes.Field(i)
		Value := InputValues.Field(i).Interface()

		fmt.Printf("%s:%v=%v\n", Field.Name, Field.Type, Value)

	}

	for i := 0; i < InputTypes.NumMethod(); i++ {
		Method := InputTypes.Method(i)
		fmt.Printf("%s: %v\n", Method.Name, Method.Type)
	}

}

func main() {

	stu := Student1{"Mhl", "girl", 24}

	GetTypeValue(stu)

}
