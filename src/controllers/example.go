package controllers

import (
	"github.com/sirupsen/logrus"
	"net/http"
	// For API only services it might
	// not be neccessary to use views.
	"views"
	// Use models to access data.
	"models"
)

func Example(w http.ResponseWriter, r *http.Request) {
	/*
		Example web page.
	*/

	// Logging the start and end of the request.
	logrus.WithFields(logrus.Fields{
		"stage":    "controller",
		"function": "Example()",
		"url":      r.URL,
		"resp":		w,
	}).Debug("Request recieved")

	defer logrus.WithFields(logrus.Fields{
		"stage":    "controller",
		"function": "Example()",
		"url":      r.URL.Path,
	}).Debug("Request processed")

	// Get data from model
	modelData := models.GetExampleData()

	// Transform this into data for view
	viewData := views.ExampleData {
		Message: modelData.Message,
	}

	// Pass the data to the view.
	views.Example(w, viewData)
}
