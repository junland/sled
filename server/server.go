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

const (
	defLvl  = "info"
	defPort = "8080"
	defPID  = "/var/run/sled.pid"
	defTLS  = "false"
)

// Sets up and starts the main server application
func Start() {
	// Get log level enviroment variable.
	envLvl, err := log.ParseLevel(utils.GetEnv("SLED_LOG_LVL", defLvl))
	if err != nil {
		fmt.Println("Invalid log level ", utils.GetEnv("SLED_LOG_LVL", defLvl))
	} else {
		// Setup logging with Logrus.
		log.SetLevel(envLvl)
	}

	envCert := utils.GetEnv("SLED_CERT", "")
	envKey := utils.GetEnv("SLED_KEY", "")
	envTLS := utils.GetEnv("SLED_TLS", defTLS)

	if envTLS == "true" {
		if envCert == "" || envKey == "" {
			log.Fatal("Invalid TLS configuration, please pass a file path for both SLED_KEY and SLED_CERT")
		}
	}

	log.Info("Setting up server...")

	envPort := utils.GetEnv("SLED_SRV_PORT", defPort)
	envPID := utils.GetEnv("SLED_PID_FILE", defPID)

	log.Debug("Setting route info...")

	// Set the router.
	router := httprouter.New()

	// Set the routes for the application.
	router.GET("/hello", helloGlobalHandle)
	router.GET("/hello/:name", helloNameHandle)

	// Chain middleware using Alice.
	chain := alice.New(SimpleLogger).Then(router)

	srv := &http.Server{Addr: ":" + envPort, Handler: chain}

	log.Debug("Starting server on port ", envPort)

	go func() {
		if defTLS == "true" {
			err := srv.ListenAndServeTLS(envCert, envKey)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	log.Info("Serving on port " + envPort + ", press CTRL + C to shutdown.")

	p := utils.NewPID(envPID)

	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT

	log.Warn("Shutting down server...")

	defer p.RemovePID()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // shut down gracefully, but wait no longer than 5 seconds before halting

	srv.Shutdown(ctx)
}
