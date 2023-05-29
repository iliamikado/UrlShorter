package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeShortURL(t *testing.T) {
	urlsStorage = NewUrlsStorage()
	defer func() {urlsStorage = nil}()

	request := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(CheckRequest)
	h.ServeHTTP(w, request)
	result := w.Result()

	assert.Equal(t, result.StatusCode, 201)

	shortURL, err := io.ReadAll(result.Body)
	require.NoError(t, err)
	err = result.Body.Close()
	require.NoError(t, err)

	assert.Equal(t, string(shortURL), "http://localhost:8080/short0")
}

func TestGetLongURLByShort(t *testing.T) {
	urlsStorage = NewUrlsStorage()
	urlsStorage.AddURL("http://google.com")
	defer func() {urlsStorage = nil}()

	request := httptest.NewRequest(http.MethodGet, "/short0", nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(CheckRequest)
	h.ServeHTTP(w, request)
	result := w.Result()

	assert.Equal(t, result.StatusCode, 307)
	err := result.Body.Close()
	require.NoError(t, err)

	longURL := result.Header.Get("Location")
	assert.Equal(t, longURL, "http://google.com")
}