package handlers

import (
	"fmt"
	"net/http"

	"github.com/achsanalfitra/Basic-Web-Application/packages/divide"
	"github.com/achsanalfitra/Basic-Web-Application/packages/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// This is a function that responds to http request
	// Combined with the HandleFunc function, this function can be called to do something on the webpage
	// The syntax is that this function takes two argument, a type of ResponseWriter and a type of Request, both from the http package

	fmt.Fprintf(w, "This is the home page")
	// Fprintf is used instead of Println because we want to write on the webpage
	// In this case, this uses the http.ResponseWriter function to write "This is the home page"
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	render.TemplateParser(w, "landing-page.html")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := divide.DivideBy(10.0, 0)
	// This function is the request handler for divideBy function
	// It uses the divideBy function and returns its result for the ResponseWriter

	if err != nil {
		fmt.Fprintf(w, fmt.Sprint(err))
	} else {
		fmt.Fprintf(w, fmt.Sprint("Your division results in: ", result))
	}
	// This is the error handling structure, a pretty standard if else
}
