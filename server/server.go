package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/junland/sled/utils"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
)

// Sets up and starts the main server application
func Start(logLvl string, srvPort string, srvPID string, srvTLS bool, srvCert string, srvKey string) {
	// Get log level enviroment variable.
	envLvl, err := log.ParseLevel(logLvl)
	if err != nil {
		fmt.Println("Invalid log level ", logLvl)
	} else {
		// Setup logging with Logrus.
		log.SetLevel(envLvl)
	}

	if srvTLS == true {
		if srvCert == "" || srvKey == "" {
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

	srv := &http.Server{Addr: ":" + srvPort, Handler: chain}

	log.Debug("Starting server on port ", srvPort)

	go func() {
		if srvTLS == true {
			err := srv.ListenAndServeTLS(srvCert, srvKey)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	log.Info("Serving on port " + srvPort + ", press CTRL + C to shutdown.")

	p := utils.NewPID(srvPID)

	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT

	log.Warn("Shutting down server...")

	defer p.RemovePID()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // shut down gracefully, but wait no longer than 5 seconds before halting

	srv.Shutdown(ctx)
}
