package main

import (
	. "github.com/velomatrix/surge/api/app/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/run/{id}", ListRunsHandler)
	http.HandleFunc("/runs", ListRunsHandler)
	http.ListenAndServe(":8080", nil)
}
