package main

import (
	"strconv"
)

type UrlsStorage struct {
	shortToLong map[string]string
	longToShort map[string]string
	count int
}

func NewUrlsStorage() *UrlsStorage {
	return &UrlsStorage{
		shortToLong: make(map[string]string),
		longToShort: make(map[string]string),
		count: 0,
	}
}

func (us *UrlsStorage) AddURL(longURL string) string {
	shortURL := "short" + strconv.Itoa(us.count)

	us.shortToLong[shortURL] = longURL
	us.longToShort[longURL] = shortURL
	us.count++;
	return shortURL
}


func (us *UrlsStorage) GetLongURLByShort(shortURL string) (string, bool) {
	longURL, ok := us.shortToLong[shortURL]
	return longURL, ok
}