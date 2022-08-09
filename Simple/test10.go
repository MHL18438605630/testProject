package main

import "fmt"

func main() {

	stringArr := map[int]string{
		1: "mhl",
		2: "xxxx",
		3: "cccc",
		4: "bbbb",
	}

	for i := 1; i <= len(stringArr); i++ {
		fmt.Println(stringArr[i])
	}

	for i := range stringArr {
		fmt.Println(stringArr[i])
	}

	for k, v := range stringArr {
		fmt.Println(k)
		fmt.Println(v)
	}

	for _, name := range stringArr {
		fmt.Println(name)
	}

}
