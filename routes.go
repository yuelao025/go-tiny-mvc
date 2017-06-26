package main

/*
	This defines the routes that the net/http
	multiplexor uses to decide how it will
	route requests to the controllers.

	Custom routes should be defined in the
	addRoute() function.
*/

import (
	"controllers"
	"github.com/sirupsen/logrus"
	"net/http"
)

//	Defines the routing decisions to be taken
//	by the HTTP server.
//
// Create the multiplexor that will
// handle routing decisions. Concurrency
// is build into the http.ServeMux type
// so goroutines and channels are not
// required in the controllers.
func addRoutes(assetDir string) *http.ServeMux {

	logrus.WithFields(logrus.Fields{
	}).Info("Starting building routing table")

	defer logrus.WithFields(logrus.Fields{
	}).Info("Finished building routing table")

	mux := http.NewServeMux()

	// Add the asset serving path
	defer logrus.WithFields(logrus.Fields{
	}).Info("Serving assets from ", assetDir)
	assetServe := "/assets/"
	mux.Handle(assetServe, http.StripPrefix(assetServe, http.FileServer(http.Dir(assetDir))))
	
	// Define custom routes here
	mux.HandleFunc("/", controllers.Example)
	// End custom routes.
	return mux
}
