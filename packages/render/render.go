package render

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/achsanalfitra/Basic-Web-Application/packages/config"
)

var app *config.AppConfig

func AccessTemplateCache(a *config.AppConfig) {
	app = a
	app.TemplateLoaded = "Template is loaded"
}

func TemplateParser(w http.ResponseWriter, html string) {
	_, cacheContent := app.TemplateCache[html]

	if !cacheContent {
		fmt.Println("Creating new template for", html)
		templateCacher(html)
	} else {
		fmt.Println("Using existing template for", html)
	}

	cachedTemplate := app.TemplateCache[html]

	err := cachedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the HTML template")
	}
}

func templateCacher(html string) {
	createdTemplate, _ := template.ParseFiles("./html-templates/" + html)
	app.TemplateCache[html] = createdTemplate
}
