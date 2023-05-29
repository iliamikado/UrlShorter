package main

import (
	"strconv"
)

type UrlsStorage struct {
	shortToLong map[string]string
	longToShort map[string]string
	shortById map[int]string
	count int
}

func NewUrlsStorage() *UrlsStorage {
	return &UrlsStorage{
		shortToLong: make(map[string]string),
		longToShort: make(map[string]string),
		shortById: map[int]string{},
		count: 0,
	}
}

func (us *UrlsStorage) AddUrl(longUrl string) string {
	shortUrl := "short" + strconv.Itoa(us.count)

	us.shortToLong[shortUrl] = longUrl
	us.longToShort[longUrl] = shortUrl
	us.shortById[us.count] = shortUrl
	us.count++;
	return shortUrl
}

func (us *UrlsStorage) GetShortUrlById(id int) (string, bool) {
	shortUrl, ok := us.shortById[id]
	return shortUrl, ok
}

func (us *UrlsStorage) GetLongUrlByShort(shortUrl string) (string, bool) {
	longUrl, ok := us.shortToLong[shortUrl]
	return longUrl, ok
}