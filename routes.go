package main

/*
	This defines the routes that the net/http
	multiplexor uses to decide how it will
	route requests to the controllers.

	Custom routes should be defined in the 
	addRoute() function.
*/

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"controllers"
)

func addRoutes() (*http.ServeMux) {
	/*
		Defines the routing decisions to be taken
		by the HTTP server.
	*/
	// Create the multiplexor that will
	// handle routing decisions. Concurrency
	// is build into the http.ServeMux type
	// so goroutines and channels are not
	// required in the controllers.
	mux := http.NewServeMux()

	// Define custom routes here
	mux.HandleFunc("/", controllers.Example)

	// End custom routes.
	logrus.WithFields(logrus.Fields{
		"mux": mux,
	}).Info("Finished building routing table.")
	return mux
}