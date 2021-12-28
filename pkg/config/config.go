package config

// You need to make sure that the configuartion files imports only the standard library packages and not other packages

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache bool
	// create a variable of type map that holds key:value pairs of name of template(key of type string) and a value which is a pointer to a template
	TemplateCache map[string]*template.Template

	// variable that can hold logs
	InfoLog *log.Logger
}
