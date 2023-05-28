package main

import (
    "net/http"
	"fmt"
)

func CheckRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		MakeShortUrl(w, r)
	case http.MethodPost:
		MakeShortUrl(w, r)
	default:
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong method")
	}
}

func MakeShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

func GetUrlById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}


func main() {
    http.HandleFunc("/", CheckRequest)
	
    http.ListenAndServe(":8080", nil)
} 