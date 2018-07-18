package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// CommonFormatPattern is the Apache Common Log format.
	CommonFormatPattern = "%s - %s [%s] \"%s\" %d %d\n"
)

// LogRequest describes a request that is made into the server.
type LogRequest struct {
	http.ResponseWriter
	Time                                            time.Time
	RemoteIP, Method, URI, Protocol, Username, Host string
	Status                                          int
	ResponseBytes                                   int
	ElapsedTime                                     time.Duration
	RequestHeader                                   http.Header
}

// AccessLogger configures a global HTTP access log.
func AccessLogger(handler http.Handler, e bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		if e != true {
			// This should exit the logging handler.
			return
		}

		sw := LogRequest{ResponseWriter: w}

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

		handler.ServeHTTP(&sw, r)

		duration := time.Now().Sub(startTime)

		record := &LogRequest{
			Time:          startTime,
			RemoteIP:      clientIP,
			Method:        r.Method,
			URI:           r.RequestURI,
			Protocol:      r.Proto,
			Username:      username,
			Host:          r.Host,
			Status:        sw.Status,
			ResponseBytes: sw.ResponseBytes,
			ElapsedTime:   duration,
		}

		requestLine := fmt.Sprintf("%s %s %s", record.Method, record.URI, record.Protocol)

		fmt.Printf(CommonFormatPattern, record.RemoteIP, record.Username, record.Time.Format("02/Jan/2006:03:04:05 -0700"), requestLine, record.Status, record.ResponseBytes)
	}
}

// WriteHeader overrides the default WriteHeader function so that you can log HTTP statues.
func (w *LogRequest) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

// Write overrides the default Write function so that AccessLogger can work correctly.
func (w *LogRequest) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.ResponseBytes += n
	return n, err
}

// SimpleMiddleware is just a example logging middleware.
func SimpleMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("This is a simple middleware.")
		h.ServeHTTP(w, r)
	})
}

// Recovery function handles the logging of panics if the web server encounters a error.
// Once the error is logged, the server will respond with a 500 error code to the client.
func Recovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("encountered a unknown error")
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
