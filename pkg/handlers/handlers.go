package handlers

import (
	"net/http"

	"github.com/AdirNoyman/MyFirstGolangWebApp/pkg/render"
)

// Handlers = controllers
// Route functions ///////////////////////////////////////////////////////////

// w = res (response)
// r = req (request)
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.html")

}
