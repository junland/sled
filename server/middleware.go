package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// HTTPLogger struct describes a HTTP request into the server.
type HTTPLogger struct {
	Time                                            time.Time
	RemoteIP, Method, URI, Protocol, Username, Host string
	Status                                          int
	ResponseBytes                                   int64
	ElapsedTime                                     time.Duration
	RequestHeader                                   http.Header
}

const (
	// CommonFormatPattern is the Apache Common Log format.
	CommonFormatPattern = "%s - %s [%s] \"%s\" %d %d\n"
)

// AccessLogger configures a global HTTP access log.
func AccessLogger(handler http.Handler, e bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Checking if not true...")
		if e != true {
			// This should exit the logging handler.
			return
		}

		clientIP := r.RemoteAddr

		if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
			clientIP = clientIP[:colon]
		}

		username := "-"
		if r.URL.User != nil {
			if name := r.URL.User.Username(); name != "" {
				username = name
			}
		}

		startTime := time.Now()

		record := &HTTPLogger{
			Time:          startTime.UTC(),
			RemoteIP:      clientIP,
			Method:        r.Method,
			URI:           r.RequestURI,
			Protocol:      r.Proto,
			Username:      username,
			Host:          r.Host,
			Status:        r.Response.StatusCode,
			ResponseBytes: r.Response.ContentLength,
			ElapsedTime:   time.Duration(0),
		}

		handler.ServeHTTP(w, r)

		finishTime := time.Now()

		record.Time = finishTime.UTC()

		record.ElapsedTime = finishTime.Sub(startTime)

		requestLine := fmt.Sprintf("%s %s %s", record.Method, record.URI, record.Protocol)

		fmt.Printf(CommonFormatPattern, record.RemoteIP, record.Username, record.Time.Format("02/Jan/2006:03:04:05 -0500"), requestLine, record.Status, record.ResponseBytes)
	})
}

// SimpleMiddleware is just a example logging middleware.
func SimpleMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("This is a simple middleware.")
		h.ServeHTTP(w, r)
	})
}
