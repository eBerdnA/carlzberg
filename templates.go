package main

import (
	"github.com/spf13/viper"
	"html/template"
	"log"
	"path"
)

var defaultTemplate *template.Template

func init() {
	log.Print("init of templates")
	if initDefaultTemplate() != nil {
		log.Fatal("Failed to initialize templates")
		return
	}
}

func initDefaultTemplate() (err error) {
	tplDirectory := viper.GetViper().GetString("templatePath")
	defaultTemplatePath := path.Join(tplDirectory, "default")
	log.Print(defaultTemplatePath)
	defaultTemplate, err = template.ParseFiles(defaultTemplatePath)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
