package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/cat/", http.HandlerFunc(cat))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "welcome to index page")
}

func dog(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "you got the dog!")
}

func cat(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "you got the cat!")
}
