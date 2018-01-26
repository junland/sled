package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Calls the function in cmd.go
	cmdExec()

	// Setup logging with Logrus.
	log.SetOutput(os.Stdout)
}

func main() {
	stopChan := make(chan os.Signal)

	log.Info("Setting server router.")

	// Set the router.
	router := httprouter.New()

	// Set the routes for the application.
	router.GET("/hello", helloGlobalHandle)
	router.GET("/hello/:name", helloNameHandle)

	log.Info("Setting server port.")

	// Set the port for a secure port, default is 443
	portNum := ":" + GetEnv("SLED_SRV_PORT", "443")

	// Chain middleware using Alice.
	app := alice.New(SimpleLogger).Then(router)

	srv := &http.Server{Addr: portNum, Handler: app}

	log.Info("Starting server on port " + GetEnv("SLED_SRV_PORT", "443"))

	go func() {
		if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
			log.Error("Listen: %s\n", err)
		}
	}()

	p := NewPID("/var/run/sled.pid")

	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT

	log.Warn("Shutting down server...")

	defer p.RemovePID()

	// Shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
}
