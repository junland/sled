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

func TestCORS(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Origin", req.URL.String())

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	})

	rr := httptest.NewRecorder()

	testHandler := CORS(handler)

	testHandler.ServeHTTP(rr, req)

	expected := "*"
	if rr.Header().Get("Access-Control-Allow-Origin") != expected {
		t.Errorf("handler returned unexpected headers: got %v want %v", rr.Header().Get("Access-Control-Allow-Origin"), expected)
	}

	expected = "POST, GET, OPTIONS, PUT, DELETE"
	if rr.Header().Get("Access-Control-Allow-Methods") != expected {
		t.Errorf("handler returned unexpected headers: got %v want %v", rr.Header().Get("Access-Control-Allow-Methods"), expected)
	}

	expected = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	if rr.Header().Get("Access-Control-Allow-Headers") != expected {
		t.Errorf("handler returned unexpected headers: got %v want %v", rr.Header().Get("Access-Control-Allow-Headers"), expected)
	}
}
