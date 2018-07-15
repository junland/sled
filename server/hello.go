package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RootHandle is a handle.
func RootHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("I am root."))
}

// helloGlobalHandle is a example handler.
func helloGlobalHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello everyone!"))
}

// helloNameHandle is a example parameter handler.
func helloNameHandle(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	name := ps.ByName("name")
	w.WriteHeader(203)
	fmt.Fprintf(w, "Hello, %s\n", name)
}
