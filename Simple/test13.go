package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

// 加密 + 生成签名

func main() {

	params := map[string]interface{}{
		"name": "MHL",
		"pvd":  "123456",
		"age":  30,
		"sex":  "girl",
	}
	fmt.Printf("sign =%s\n", createSign(params))

}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func createSign(params map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)

	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str += fmt.Sprintf("&xl_%v=%v", key[i], params[key[i]])
		}
	}

	var secret = "123456789"

	sign := MD5(MD5(str) + MD5(secret))
	return sign

}

