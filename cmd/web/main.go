package main

import (
	// "errors"
	"fmt"
	"log"

	"net/http"

	"github.com/AdirNoyman/MyFirstGolangWebApp/pkg/config"
	"github.com/AdirNoyman/MyFirstGolangWebApp/pkg/handlers"
	"github.com/AdirNoyman/MyFirstGolangWebApp/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig
	// tc = template cache
	tc, err := render.CreateTemplatesCache()

	if err != nil {

		log.Fatal("Can't create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// server //////////////////////////////////////////////////////
	// this function returns an error. In this case we decided not to store it, so that's why we assigned it to nothing ('_')
	fmt.Printf("Server started and listening on port %s 😎🤘", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
