package main

import (
	// "errors"
	"fmt"

	"net/http"

	"github.com/AdirNoyman/MyFirstGolangWebApp/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// server //////////////////////////////////////////////////////
	// this function returns an error. In this case we decided not to store it, so that's why we assigned it to nothing ('_')
	fmt.Printf("Server started and listening on port %s 😎🤘", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
