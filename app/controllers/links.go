package controllers

import (
	"github.com/revel/revel"
)

type Links struct {
	*revel.Controller
}

func (c Links) List() revel.Result {
	return c.Render()
}
