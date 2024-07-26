package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func TemplateParser(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./html-templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the HTML template")
	}
}
