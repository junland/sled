package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// helloGlobalHandle is a example handler.
func helloGlobalHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Everyone!\n")
}

// helloNameHandle is a example paramter handler.
func helloNameHandle(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	name := ps.ByName("name")
	fmt.Fprintf(w, "Hello, %s\n", name)
}
