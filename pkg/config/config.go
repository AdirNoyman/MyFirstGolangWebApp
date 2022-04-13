package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	// UseCache - If we are in production will use the template cache. But if we are in development, we will re-build it every time we load the app
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
