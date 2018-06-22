package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SimpleLogger is just a example logging middleware.
func SimpleLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.URL, " requested from ", r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}
