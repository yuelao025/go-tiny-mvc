// This package holds all of the rendered HTML that the
// end user sees. Data is passed to it via API called
// to the controller.
package views

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type ExampleData struct {
	Message string
}

// Example function to show how a view should be implented.
func Example(w http.ResponseWriter, data ExampleData) error {
	var err error

	// Logging the start and end of the request.
	logrus.WithFields(logrus.Fields{
		"stage": "view",
		"data":  data,
	}).Debug("Example view triggered")

	defer logrus.WithFields(logrus.Fields{
		"stage": "view",
		"data":  data,
	}).Debug("Example view processed")

	// Add the <head> tags.
	h := head {}
	h.setTitle("Example page")
	err = generateHead(w, h)

	// Start templating.
	t := template.New("example")
	tp, err := t.Parse(`
<html>
{{.Message}}
<br>
<img src="/assets/images/penguin.jpg" alt="Penguin"> 
</html>
`)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"stage":    "view",
			"function": "Example()",
			"data":     data,
			"template": t,
		}).Debug("Could not render view.")
	}

	tp.Execute(w, data)
	return err
}
