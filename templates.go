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
	defaultTemplate, err = template.New("List").Parse(
		"<!doctype html>" +
			"<html lang=en>" +
			"<meta name=viewport content=\"width=device-width, initial-scale=1.0\">" +
			"<title>Carlzberg</title>" +
			"<h1>Hello World from Carlzberg</h1>" +
			"</html>")
	return
}
