package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccessLogger(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	})

	rr := httptest.NewRecorder()

	logger := AccessLogger(handler, true)

	logger.ServeHTTP(rr, req)
}

func TestRecovery(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		panic("This is going to blow up.")
	})

	rr := httptest.NewRecorder()

	testHandler := Recovery(handler)

	testHandler.ServeHTTP(rr, req)

	expected := http.StatusInternalServerError
	if rr.Code != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Code, expected)
	}
}
