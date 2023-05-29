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

func (us *UrlsStorage) AddUrl(longUrl string) string {
	shortUrl := "short" + strconv.Itoa(us.count)

	us.shortToLong[shortUrl] = longUrl
	us.longToShort[longUrl] = shortUrl
	us.count++;
	return shortUrl
}


func (us *UrlsStorage) GetLongUrlByShort(shortUrl string) (string, bool) {
	longUrl, ok := us.shortToLong[shortUrl]
	return longUrl, ok
}