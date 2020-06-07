package controllers

import (
	"github.com/revel/revel"
)

type Links struct {
	*revel.Controller
}

func (c Links) Listing() revel.Result {
	return c.Render()
}
