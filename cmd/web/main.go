package main

import (
	"fmt"
	"net/http"

	"github.com/achsanalfitra/Basic-Web-Application/packages/handlers"
)

const port string = "8080"
const homeAddress string = "/"
const divideAddress string = "/divide"
const landingPageAddress string = "/landing-page"

// Create the port and homeAddress as a constant because they are not going to be changed
// Also, these parameters are not changing so its good to make them constant

func main() {
	http.HandleFunc(homeAddress, handlers.Home)
	// On the main function, the HandleFunc function takes the request responder function Home that has been made previously

	http.HandleFunc(divideAddress, handlers.Divide)
	// This creates another web page on the divisionAddress of the localhost:8080 port
	// It runs the Divide function

	http.HandleFunc(landingPageAddress, handlers.LandingPage)

	fmt.Println("Running web application on port:", port)
	// This is just to create a message that a web application is running on certain port locally

	_ = http.ListenAndServe(fmt.Sprint(":", port), nil)
	// This function runs the web application and listen to every request
	// I honestly don't understand why assigning this to a value actually starts it too
	// The assignment is underscored because this function returns an error and we do not care about that error
}
