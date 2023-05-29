package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func CheckRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetURLById(w, r)
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
	shortURL := urlsStorage.AddURL(string(body))
	w.WriteHeader(201)
	w.Write([]byte(shortURL))
}

func GetURLById(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")

	longURL, ok := urlsStorage.GetLongURLByShort(shortURL)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong short URL")
		return
	}

	w.WriteHeader(307)
	w.Header().Add("Location", longURL)
	w.Write([]byte(longURL))
}

var urlsStorage *UrlsStorage;

func main() {
	urlsStorage = NewUrlsStorage()

    http.HandleFunc("/", CheckRequest)
    http.ListenAndServe(":8080", nil)
} 