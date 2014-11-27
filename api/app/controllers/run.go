package controllers

import (
	"github.com/revel/revel"
	"github.com/velomatrix/surge/api/app/models"
)

type Run struct {
	*revel.Controller
}

// temporary data store
var runs = [6]*models.Run{
	&models.Run{RunId: 0, Name: "Run 0"},
	&models.Run{RunId: 1, Name: "Run 1"},
	&models.Run{RunId: 2, Name: "Run 2"},
	&models.Run{RunId: 3, Name: "Run 3"},
	&models.Run{RunId: 4, Name: "Run 4"},
	&models.Run{RunId: 5, Name: "Run 5"},
}

// returns a list of runs
func (c Run) Index() revel.Result {
	return c.RenderJson(runs)
}

// returns a single run
func (c Run) Show(id int) revel.Result {
	return c.RenderJson(runs[id])
}