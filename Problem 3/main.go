package main

import (
	"Human"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/nama", Human.HumanData)
	mux.HandleFunc("/semuadata", Human.AllHumanData)

	http.ListenAndServe(":5050", mux)
}
