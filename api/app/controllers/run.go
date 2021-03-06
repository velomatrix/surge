package controllers

import (
	"github.com/revel/revel"
	"github.com/velomatrix/revmgo"
	"github.com/velomatrix/surge/api/app/services"
)

type Run struct {
	*revel.Controller
	revmgo.MongoController
}

// returns a list of runs
func (c Run) Index() revel.Result {
	runs, _ := services.FindAllRuns()
	return c.RenderJson(runs)
}

// returns a single run
func (c Run) Show(id string) revel.Result {
	run, _ := services.FindRun(id)
	return c.RenderJson(run)
}
