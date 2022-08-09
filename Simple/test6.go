package main

import (
	"encoding/json"
	"fmt"
)

type ResultMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	var res ResultMessage
	res.Code = 200
	res.Message = "success"

	toJson(&res)
	setMessage(&res)
	toJson(&res)
}

func setMessage(res *ResultMessage) {
	res.Code = 500
	res.Message = "fail"
}

func toJson(message *ResultMessage) {
	json, errs := json.Marshal(message)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data:", json)
}
