package main

import (
	"html/template"
	"log"
)

var defaultTemplate *template.Template

func init() {
	if initDefaultTemplate() != nil {
		log.Fatal("Failed to initialize templates")
		return
	}
}

func initDefaultTemplate() (err error) {
	defaultTemplate, err = template.New("Post").Parse(
		"<!doctype html>" +
			"<html lang=en>" +
			"<meta name=viewport content=\"width=device-width, initial-scale=1.0\">" +
			"<title>{{ .Post_title}}</title>" +
			"<h1>{{ .Post_content}}</h1>" +
			"</html>")
	return
}
