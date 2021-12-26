package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// decalring a variable that holds a map of functions. name of function as key and value is the function itself
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)

	if err != nil {

		fmt.Println("Error getting template cache: ", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)

	if err != nil {

		fmt.Println("error parsing template: ", err)
		return
	}

}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

	// create a map (key:value data structure) that will hold all my templates. the key will be the template name and the value is the parsed template
	myCache := map[string]*template.Template{}

	// get all files that are in the templates folder, that there name is *[some name].page.html
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {

		return myCache, err
	}

	// the first parameter is the index, which in this case we decided not to use it and thts why we are using here the sign '_'
	// pages = all the files we found in the templates folder
	for _, page := range pages {

		fmt.Println("Page is currently", page)
		// assigning the file name to name variable
		name := filepath.Base(page)

		// creating a set of templates
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {

			return myCache, err
		}

		// check if we have any matching layout that suppose to be in this template
		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {

			return myCache, err
		}

		// if we find any layout match, than the length of 'matches' variable will be equal to 1 or more
		if len(matches) > 0 {

			// parse the matching files
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {

				return myCache, err

			}

			myCache[name] = ts
		}

	}
	return myCache, nil
}
