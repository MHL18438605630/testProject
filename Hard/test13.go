package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", writeIntoDB(greet))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

// DB

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,%q", html.EscapeString(r.URL.Path))

}

func writeIntoDB(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// 进行其别操作

		fmt.Println("write to db success")
		f(w, r)
	}

}
