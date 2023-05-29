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
		GetUrlById(w, r)
	case http.MethodPost:
		MakeShortUrl(w, r)
	default:
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong method")
	}
}

func MakeShortUrl(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong body")
		return
	}
	shortUrl := urlsStorage.AddUrl(string(body))
	w.WriteHeader(201)
	w.Write([]byte(shortUrl))
}

func GetUrlById(w http.ResponseWriter, r *http.Request) {
	shortUrl := strings.TrimPrefix(r.URL.Path, "/")

	longUrl, ok := urlsStorage.GetLongUrlByShort(shortUrl)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Wrong short url")
		return
	}

	w.WriteHeader(307)
	w.Header().Add("Location", longUrl)
	w.Write([]byte(longUrl))
}

var urlsStorage *UrlsStorage;

func main() {
	urlsStorage = NewUrlsStorage()

    http.HandleFunc("/", CheckRequest)
    http.ListenAndServe(":8080", nil)
} 