package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func helloGlobalHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello Everyone!\n")
}

func helloNameHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}
