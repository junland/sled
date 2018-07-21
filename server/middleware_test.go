package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}

func TestAccessLogger(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	})

	logger := AccessLogger(handler, true)
	logger.ServeHTTP(httptest.NewRecorder(), newRequest("GET", "/"))
}
