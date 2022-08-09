package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			goto end
		}
		fmt.Println(i)
	}
end:
	fmt.Println("endddddd")
}
