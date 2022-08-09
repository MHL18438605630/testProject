package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"price"`
	Actor []string `json:"actor"`
}

func main() {

	mov := Movie{"花样年华", 2000, 35, []string{"李建", "孙俪"}}
	JsonStr, err1 := json.Marshal(mov)
	if err1 != nil {
		fmt.Println("编码错误")
		return
	}
	fmt.Printf("JsonStr =%s \n", JsonStr)

	MyMov := Movie{}

	err1 = json.Unmarshal(JsonStr, &MyMov)

	if err1 != nil {
		fmt.Println("解码失败")
		return
	}

	fmt.Printf("MyMov =%v \n", MyMov)

}
