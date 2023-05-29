package main

import (
	"fmt"
	"io"
	"net/http"
)

func CheckRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetLongURLByShort(w, r)
	case http.MethodPost:
		MakeShortURL(w, r)
	default:
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong method")
	}
}

func MakeShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong body")
		return
	}
	shortURL := host + urlsStorage.AddURL(string(body))
	w.WriteHeader(201)
	w.Write([]byte(shortURL))
}

func GetLongURLByShort(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path

	longURL, ok := urlsStorage.GetLongURLByShort(shortURL)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong short URL")
		return
	}

	w.WriteHeader(307)
	w.Header().Set("Location", longURL)
	w.Write([]byte(longURL))
}

var urlsStorage *UrlsStorage
var host = "http://localhost:8080"

func main() {
	urlsStorage = NewUrlsStorage()

    http.HandleFunc("/", CheckRequest)
    http.ListenAndServe(":8080", nil)
} 