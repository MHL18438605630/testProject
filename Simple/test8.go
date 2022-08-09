package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	res["data"] = map[string]interface{}{
		"username": "MHL",
		"age":      24,
		"hobby":    []string{"读书", "喜欢BL剧"},
	}
	fmt.Println(res)

	// 序列化
	jsons, errs := json.Marshal(res)

	if errs != nil {
		fmt.Println("json marshal error", errs)
	}
	fmt.Println(jsons)

	// 反序列化

	var res2 map[string]interface{}
	errs2 := json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("errors", errs2)
	}
	fmt.Println(res2)

}
