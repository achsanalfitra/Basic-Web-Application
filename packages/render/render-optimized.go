package render

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func OptRenderTemplate(w http.ResponseWriter, html string) {
	// create a template cache
	// get the requested template from cache
	// render the template
}

func createRenderTemplate() (map[string]*template.Template, error) {
	var templateCache map[string]*template.Template

	// get all the template files from ./html-templates
	htmlPages, err := filepath.Glob("./html-templates/*-page.html")
	if err != nil {
		return templateCache, err
	}

	// range through all html pages
	for _, page := range htmlPages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		templateCache[name] = templateSet
	}

	return templateCache, err

}
