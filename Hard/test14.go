package main

import (
	"fmt"
	"net/http"
)

func main() {
	db := NewDatabase("localhost:5432")
	http.HandleFunc("/hello", Hello(db))
	http.ListenAndServe(":3000", nil)
}

type Database struct {
	Url string
}

func NewDatabase(url string) Database {
	return Database{url}
}
func Hello(db Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.Url)
	}
}
