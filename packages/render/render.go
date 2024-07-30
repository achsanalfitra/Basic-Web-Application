package render

import (
	"fmt"
	"html/template"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func TemplateParser(w http.ResponseWriter, html string) {
	_, cacheContent := templateCache[html]

	if !cacheContent {
		fmt.Println("Creating new template for", html)
		templateCacher(html)
	} else {
		fmt.Println("Using existing template for", html)
	}

	cachedTemplate := templateCache[html]

	err := cachedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the HTML template")
	}
}

func templateCacher(html string) {
	createdTemplate, _ := template.ParseFiles("./html-templates/" + html)
	templateCache[html] = createdTemplate
}
