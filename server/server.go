package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

// Config struct provides configuration fields for the server.
type Config struct {
	LogLvl string
	Port   string
	PID    string
	TLS    bool
	Cert   string
	Key    string
}

var done = make(chan os.Signal)

// Start sets up and starts the main server application
func Start(c Config) error {
	// Get log level environment variable.
	envLvl, err := log.ParseLevel(c.LogLvl)
	if err != nil {
		fmt.Println("Invalid log level ", envLvl)
	} else {
		// Setup logging with Logrus.
		log.SetLevel(envLvl)
	}

	if c.TLS == true {
		if c.Cert == "" || c.Key == "" {
			log.Fatal("Invalid TLS configuration, please pass a file path for both SLED_KEY and SLED_CERT")
		}
	}

	log.Info("Setting up server...")

	router := RegisterRoutes()

	log.Debug("Setting up logging...")

	value := true

	srv := &http.Server{Addr: ":" + c.Port, Handler: AccessLogger(router, value)}

	log.Debug("Starting server on port ", c.Port)

	go func() {
		if c.TLS == true {
			err := srv.ListenAndServeTLS(c.Cert, c.Key)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	log.Info("Serving on port " + c.Port + ", press CTRL + C to shutdown.")

	p := NewPID(c.PID)

	signal.Notify(done, os.Interrupt)

	log.Warn("After notify...")

	<-done // wait for SIGINT

	log.Warn("Shutting down server...")

	p.RemovePID()

	log.Debug("Shutting down server...gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second) // shut down gracefully, but wait no longer than 45 seconds before halting

	log.Debug("Shutting down server...gracefully...even more...")

	defer cancel()

	log.Debug("Shutting down server...gracefully...even more...For sure...")

	srv.Shutdown(ctx)

	log.Debug("JK...this is actually the end of this function..")

	return nil
}

//func Stop() {
//	log.Warn("Stopping server")
//  done <- os.Interrupt
//}
