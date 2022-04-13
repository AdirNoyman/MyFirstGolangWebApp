package main

import (
	"fmt"
	"hello_world3/pkg/config"
	"hello_world3/pkg/handlers"
	"hello_world3/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache 😩")
	}

	app.TemplateCache = tc

	// Because I'm in development mode, I'm re-building the templates every time the page reloads
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// Routes
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	// Server
	fmt.Println(fmt.Sprintf("Starting application on port %s 😎🤟", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
