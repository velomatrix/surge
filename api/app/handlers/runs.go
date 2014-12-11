package handlers

import (
	"net/http"
	/*
		"encoding/json"
		"github.com/velomatrix/surge/api/app/services"
		"gopkg.in/mgo.v2/bson"
	*/)

func ShowRunHandler(w http.ResponseWriter, r *http.Request) {
	/*
		runId := r.URL.Query().Get(":runId")
		run, _ := services.FindRun(runId)
		json.NewEncoder(w).Encode(run)
	*/
}

func ListRunsHandler(w http.ResponseWriter, r *http.Request) {
	/*
		runs, _ := services.FindAllRuns()
		json.NewEncoder(w).Encode(runs)
	*/
}
