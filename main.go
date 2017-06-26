package main

// Imports

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Constants

// Change the NAME and VERSION constants to be
// to your project.
const (
	NAME    = "YOUR APP NAME HERE"
	VERSION = "0.1.0"

	ERROR_CONFIGURATION = 10
	ERROR_ROUTING       = 11
	ERROR_SERVER_DIED   = 11
)

// Globals

var configPath string = fmt.Sprintf("/etc/%s.yaml", NAME)

// Functions

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
	//"animal": "walrus",
	//"size":   10,
	}).Info("Starting a new run of ", NAME, " ", VERSION)
	parseArgs()
}

func parseArgs() {
	flag.StringVar(&configPath, "config-file", configPath, "Path to configration file.")
	flag.Parse()
}

func main() {
	// Extract configuration.
	serverConf, err := configure(configPath)

	bind := fmt.Sprintf("%s:%d", serverConf.Address, serverConf.Port)
	assetDir := serverConf.AssetDirectory

	logrus.WithFields(logrus.Fields{
		"serverConfig": serverConf,
	}).Debug("Server configuration loaded")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":     err,
			"errorcode": ERROR_CONFIGURATION,
		}).Error("Could not extract the configuration")
		os.Exit(ERROR_CONFIGURATION)
	}

	logrus.WithFields(logrus.Fields{}).Debug("Configuration loaded successfully")

	// Load routing definitions
	logrus.WithFields(logrus.Fields{}).Debug("Loading routes")

	// Load the custom routes
	mux := addRoutes(assetDir)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":     err,
			"errorcode": ERROR_ROUTING,
		}).Error("Could not load routes")
		os.Exit(ERROR_CONFIGURATION)
	}

	logrus.WithFields(logrus.Fields{}).Debug("Routes loaded successfully")

	logrus.WithFields(logrus.Fields{
		"bind": bind,
	}).Info("Starting server")
	err = http.ListenAndServe(bind, mux)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"bind":      bind,
			"errorcode": ERROR_SERVER_DIED,
			"error":     err,
		}).Info("Starting server")
		os.Exit(ERROR_SERVER_DIED)
	}
}
