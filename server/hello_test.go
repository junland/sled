package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHelloRootHandle(t *testing.T) {
	// Setup handler
	handler := http.HandlerFunc(helloRootHandle)

	// Set up the request.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up the testing recorder.
	rr := httptest.NewRecorder()

	// Set up the testing recorder.
	handler.ServeHTTP(rr, req)

	// Check if status is correct.
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

	// Check if the response was OK.
	expected := "I am root."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHelloGlobalHandle(t *testing.T) {
	// Setup handler
	handler := http.HandlerFunc(helloGlobalHandle)

	// Set up the request.
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up the testing recorder.
	rr := httptest.NewRecorder()

	// Set up the testing recorder.
	handler.ServeHTTP(rr, req)

	// Check if status is correct.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response was OK.
	expected := "Hello everyone!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHelloNameHandle(t *testing.T) {
	// Set up the router.
	router := httprouter.New()
	router.Handler("GET", "/hello/:name", http.HandlerFunc(helloNameHandle))

	// Set up the request.
	req, err := http.NewRequest("GET", "/hello/john", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up the testing recorder.
	rr := httptest.NewRecorder()

	// Send the request.
	router.ServeHTTP(rr, req)

	// Check if status is correct.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check if the response was OK.
	expected := "Hello, john\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
