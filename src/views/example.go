package views

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"io"
)

type ExampleData struct {
	Message string
}

func Example(w http.ResponseWriter, data ExampleData) {
	/*
		Example web page.
	*/

	// Logging the start and end of the request.
	logrus.WithFields(logrus.Fields{
		"stage": "view",
		"function": "Example()",
		"data": data,
	}).Debug("View triggered")

	defer logrus.WithFields(logrus.Fields{
		"stage": "view",
		"function": "Example()",
		"data": data,
	}).Debug("View processed")

	// Pass the data to the view.
	io.WriteString(w, data.Message)
}