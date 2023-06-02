package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	urlsStorage = NewUrlsStorage()
    r := NewRouter()
    ts := httptest.NewServer(r)
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 0 {
			return errors.New("Остановлено после двух Redirect")
		}
		return nil
	}

    defer ts.Close()
	var (
		statusCode int
		body string
		header http.Header
	)

	statusCode, body, _ = testRequest(t, ts, "POST", "/", strings.NewReader("http://google.com"))
    assert.Equal(t, 201, statusCode)
	body = strings.Split(body, "/")[len(strings.Split(body, "/")) - 1]

    statusCode, body, header = testRequest(t, ts, "GET", "/" + body, nil)
    assert.Equal(t, 307, statusCode)
    assert.Equal(t, "http://google.com", header.Get("Location"))
	assert.Equal(t, "Success request", body)

	statusCode, _, _ = testRequest(t, ts, "GET", "/abracadabra", nil)
	assert.Equal(t, 400, statusCode)
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (int, string, http.Header) {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

    req, err := http.NewRequest(method, ts.URL+path, body)
    require.NoError(t, err)

    resp, err := client.Do(req)
    require.NoError(t, err)

    respBody, err := io.ReadAll(resp.Body)
    require.NoError(t, err)

    defer resp.Body.Close()

    return resp.StatusCode, string(respBody), resp.Header
} 