package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func helloGlobalHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello Everyone!\n")
}

func helloNameHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}
