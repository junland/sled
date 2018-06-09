package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/junland/sled/utils"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	LogLvl string
	Port   string
	PID    string
	TLS    bool
	Cert   string
	Key    string
}

// Sets up and starts the main server application
func Start(c Config) {
	// Get log level enviroment variable.
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

	log.Debug("Setting route info...")

	// Set the router.
	router := httprouter.New()

	// Set the routes for the application.
	router.GET("/hello", helloGlobalHandle)
	router.GET("/hello/:name", helloNameHandle)

	// Chain middleware using Alice.
	chain := alice.New(SimpleLogger).Then(router)

	srv := &http.Server{Addr: ":" + c.Port, Handler: chain}

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

	p := utils.NewPID(c.PID)

	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT

	log.Warn("Shutting down server...")

	defer p.RemovePID()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // shut down gracefully, but wait no longer than 5 seconds before halting

	srv.Shutdown(ctx)
}
