package controllers

import (
	"github.com/revel/revel"
)

type Check struct {
	*revel.Controller
}

// server healthcheck
func (c Check) Health() revel.Result {
	return c.RenderText("Healthy")
}
