package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat/", cat)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to index page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "you got the dog!")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "you got the cat!")
}
