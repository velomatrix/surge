package main

import (
	. "github.com/velomatrix/surge/api/app/handlers"
	"github.com/velomatrix/surge/api/app/lib"
	"net/http"
)

func main() {
	mongo.InitDB()
	http.HandleFunc("/run/{runId}", ListRunsHandler)
	http.HandleFunc("/runs", ListRunsHandler)
	http.ListenAndServe(":8080", nil)
}
