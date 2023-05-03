package main

import (
	"Human"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //creates a new HTTP request multiplexer, called mux, which is used to route incoming requests to the appropriate handler

	mux.HandleFunc("/nama", Human.HumanData)         //registers a new request handler function for the route "/nama" that calls the HumanData function in the "Human" package when a request is made to that route
	mux.HandleFunc("/semuadata", Human.AllHumanData) // registers a new request handler function for the route "/semuadata" that calls the AllHumanData function in the "Human" package when a request is made to that route

	http.ListenAndServe(":5050", mux) //starts the web server on port 5-5-. it uses the "mux" multiplexer to route incoming req to the appropriate handler
}
