package controllers

import (
	"github.com/revel/revel"
	"github.com/velomatrix/surge/api/app/models"
)

type Run struct {
	*revel.Controller
}

func (c Run) Index() revel.Result {

	//	runs := []*models.Run{
	//		&models.Run{run_id: 0, name: "Big Load"},
	//	}

	run := models.Run{RunId: 0, Name: "Foo"}

	return c.RenderJson(run)
}
