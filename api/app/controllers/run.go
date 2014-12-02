package controllers

import (
	"github.com/revel/revel"
	"github.com/velomatrix/surge/api/app/models"
	 "gopkg.in/mgo.v2/bson"
)

type Run struct {
	*revel.Controller
}

// returns a list of runs
func (c Run) Index() revel.Result {
	runs := models.FindAllRuns()
	return c.RenderJson(runs)
}

// returns a single run
func (c Run) Show(id string) revel.Result {
	run := models.FindRun(bson.ObjectIdHex(id))
	return c.RenderJson(run)
}
