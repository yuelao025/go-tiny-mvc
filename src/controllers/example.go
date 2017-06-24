package controllers

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"views"
	//"models"
)

func Example(w http.ResponseWriter, r *http.Request) {
	/*
		Example web page.
	*/

	// Logging the start and end of the request.
	logrus.WithFields(logrus.Fields{
		"stage": "controller",
		"function": "Example()",
		"url": r.URL.Path,
	}).Debug("Request recieved")

	defer logrus.WithFields(logrus.Fields{
		"stage": "controller",
		"function": "Example()",
		"url": r.URL.Path,
	}).Debug("Request processed")

	data := views.ExampleData {
		Message: "Hello world!\n",
	}

	// Pass the data to the view.
	views.Example(w, data)
}