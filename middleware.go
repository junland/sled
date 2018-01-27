package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SimpleLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}
