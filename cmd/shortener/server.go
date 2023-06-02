package main

import (
    "net/http"
	"fmt"
	"io"
    "github.com/go-chi/chi/v5"
)

var urlsStorage *UrlsStorage
var host = "http://localhost:8080"

func NewRouter() chi.Router {
	r := chi.NewRouter()

    r.Get("/{shortURL}", func(rw http.ResponseWriter, r *http.Request) {
        shortURL := chi.URLParam(r, "shortURL")

		longURL, ok := urlsStorage.GetLongURLByShort(shortURL)
		if !ok {
			rw.WriteHeader(400)
			fmt.Fprintln(rw, "Wrong short URL")
			return
		}
	
		rw.Header().Set("Location", longURL)
		rw.WriteHeader(307)
		rw.Write([]byte("Success request"))
    })

	r.Post("/", func(rw http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			rw.WriteHeader(400)
			fmt.Fprintln(rw, "Wrong body")
			return
		}
		shortURL := host + "/" + urlsStorage.AddURL(string(body))
		rw.WriteHeader(201)
		rw.Write([]byte(shortURL))
	})
	return r
}

func StartServer() {

	urlsStorage = NewUrlsStorage()
	r := NewRouter()
    http.ListenAndServe(":8080", r)
} 