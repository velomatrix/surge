package handlers

import (
	"fmt"
	"net/http"
)

func ShowRunHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Runs!")
}

func ListRunsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All the runs")
}
