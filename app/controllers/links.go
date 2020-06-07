package controllers

import (
	"exercicio-golang-03/app"

	"github.com/didi/gendry/scanner"
	"github.com/revel/revel"
)

type Links struct {
	*revel.Controller
}

func (c Links) List() revel.Result {
	sql := "SELECT id, name, link from Link"
	c.Request.Context()

	rows, err := app.DB.Query(sql)
	if err != nil {
		return c.RenderError(err)
	}
	defer rows.Close()

	links, err := scanner.ScanMap(rows)
	if err != nil {
		return c.RenderError(err)
	}
	return c.Render(links)
}
