package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
)

// RegisterRoutes sets all the configured routes for the server to the designated handler and middleware.
func RegisterRoutes() *httprouter.Router {
	log.Debug("Setting route info...")

	// Set the router.
	router := httprouter.New()

	chain := alice.New(SimpleMiddleware)

	// Set the routes for the application.
	router.Handler("GET", "/", chain.ThenFunc(RootHandle))
	router.Handler("GET", "/hello", chain.ThenFunc(helloGlobalHandle))
	router.Handler("GET", "/hello/:name", chain.ThenFunc(helloNameHandle))

	return router
}
