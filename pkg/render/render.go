package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// decalring a variable that holds a map of functions. name of function as key and value is the function itself
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// get the template cache from the app.configuartion file
	templateCache, err := CreateTemplatesCache()

	if err != nil {

		// killing the application
		log.Fatal(err)
	}

	// if found a template will return true(ok) otherwise it will return false
	t, ok := templateCache[tmpl]

	// if didn't find a template
	if !ok {

		// killing the application
		log.Fatal(err)

	}

	// if found a template, create a bytes buffer that will hold the parsed template data(which is in bytes) and asiggn the template data to it
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	// write the template data to the response writer
	_, err = buf.WriteTo(w)

	if err != nil {

		fmt.Println("Error writing template to browser: ", err)

	}

}

// CreateTemplatesCache creates a cache of templates as map (the equivlant of dictionary/object in other languages)
func CreateTemplatesCache() (map[string]*template.Template, error) {

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
		fmt.Println("name is currently", name)
		// creating a set of templates ('ts')
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

			// parse the matching files and combines it with the template that is using this layout file
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {

				return myCache, err

			}

			myCache[name] = ts
			fmt.Println("mycache is ", myCache)
		}

	}
	return myCache, nil
}
