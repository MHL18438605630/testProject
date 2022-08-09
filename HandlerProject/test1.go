package main

import (
	"fmt"
	"net/http"
)

type Serve struct {
	ServeIP   string
	ServePort int
}

func (s *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serve: { ", s.ServeIP, ", ", s.ServePort)
}

func main() {

	http.Handle("/hello", &Serve{"127.0.0.1", 8888})
	http.ListenAndServe("127.0.0.1", nil)

}
