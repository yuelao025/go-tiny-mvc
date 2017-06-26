// Define the items in the <head> tags
// This controls all of the data that is put
// inside the <head> tags at the top of each
// page.
package views

import (
	"github.com/sirupsen/logrus"
	"text/template"
	"net/http"
)

// Contains all the information for generating <head>
type head struct {
	Title string
	Tags []tag
}

// Item to be added in a tag in <head>
type tag struct {
	// The name of the tag
	Name string
	// The values to be added into the tag
	// Example:
	// <meta key1="value2" key2="value2">
	Fields map[string]string
}

// Define common tags in <head>
// Add common <head> tags in here.
// Define each new tag inside a
// {} block to reset the scope.
func (h *head) addCommonHeads() {
	{
		tagName := "meta"
		tagValues := map[string]string {
			"test1": "value1",
		}
		h.addTag(tagName,tagValues)	
	}
	{
		tagName := "meta"
		tagValues := map[string]string {
			"http-equiv": "Content-Type",
			"content": "text/html; charset=utf-8",
		}
		h.addTag(tagName,tagValues)	
	}
}

// Adds data to be put in the <tag> tags in
// the <head> section.
// For example, to add 
//
// <meta data="value">
// meta data Add("meta", {"data": "value"})
func (h *head) addTag(tagName string, values map[string]string) {
	newTag := tag {
		Name: tagName,
		Fields: values,
	}
	h.Tags = append(h.Tags, newTag)
}

// Adds data to be put in the <title> tags in
// the <head> section.
func (h *head) setTitle(title string) {
	h.Title = title
}

// Writes the data formed in h to the <head> tags.
func generateHead(w http.ResponseWriter, h head) error {
	// Logging the start and end of the request.
	logrus.WithFields(logrus.Fields{
		"stage": "view",
	}).Debug("Building header section")

	defer logrus.WithFields(logrus.Fields{
		"stage": "view",
	}).Debug("header section built")

	// Add common head tags.
	h.addCommonHeads()

	// Start templating.
	t := template.New("example")
	tp, err := t.Parse(`<!DOCTYPE html>
<head>
	<title>{{.Title}}</title>
{{/* Iterate over all the tags and extract the key value pairs. */}}
{{range .Tags}}
	<{{.Name}} {{range $key, $value := .Fields}}{{$key}}="{{$value}}" {{- end}}>
{{end}}
</head>`)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"stage":    "view",
			"template": t,
		}).Debug("Could not render view.")
	}

	err = tp.Execute(w, h)
	return err
}