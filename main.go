package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
)

const (
	defLvl  = "info"
	defPort = "443"
)

func init() {
	// Calls the function in cmd.go
	cmdExec()

	// Gets the log level enviroment variable.
	envLvl, err := log.ParseLevel(GetEnv("SLED_LOG_LVL", defLvl))
	if err != nil {
		fmt.Println("Invalid log level %s", GetEnv("SLED_LOG_LVL", defLvl))
		os.Exit(3)
	}

	// Setup logging with Logrus.
	log.SetOutput(os.Stdout)
	log.SetLevel(envLvl)
}

func main() {
	// Gets and stores port number enviroment variable.
	envPort := GetEnv("SLED_SRV_PORT", defPort)

	log.Info("Setting server router.")

	// Set the router.
	router := httprouter.New()

	// Set the routes for the application.
	router.GET("/hello", helloGlobalHandle)
	router.GET("/hello/:name", helloNameHandle)

	log.Info("Setting server port.")

	// Set the port for a secure port, default is 443
	envport := ":" + envPort

	// Chain middleware using Alice.
	app := alice.New(SimpleLogger).Then(router)

	srv := &http.Server{Addr: envport, Handler: app}

	log.Info("Starting server on port " + envPort)

	go func() {
		if err := srv.ListenAndServeTLS("./cert.pem", "./key.pem"); err != nil {
			log.Error("Listen: %s\n", err)
		}
	}()

	p := NewPID("/var/run/sled.pid")

	// Sets gracefull shutdown.
	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT

	log.Warn("Shutting down server...")

	defer p.RemovePID()

	// Shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
}
