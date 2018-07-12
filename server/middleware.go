package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// AccessLogger configures a global HTTP access log.
func AccessLogger(handler http.Handler, v string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Here is a value: ", v)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// SimpleMiddleware is just a example logging middleware.
func SimpleMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("This is a simple middleware.")
		h.ServeHTTP(w, r)
	})
}
