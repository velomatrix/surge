package handlers

import (
	"encoding/json"
	"github.com/velomatrix/surge/api/app/services"
	"net/http"
)

func ShowRunHandler(w http.ResponseWriter, r *http.Request) {
	runId := r.URL.Query().Get(":runId")
	run, _ := services.FindRun(runId)
	json.NewEncoder(w).Encode(run)
}

func ListRunsHandler(w http.ResponseWriter, r *http.Request) {
	runs, _ := services.FindAllRuns()
	json.NewEncoder(w).Encode(runs)
}
