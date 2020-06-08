package controllers

import (
	"exercicio-golang-03/app"
	"net/http"

	"github.com/didi/gendry/scanner"
	"github.com/revel/revel"
)

type Links struct {
	*revel.Controller
}

func (c Links) List() revel.Result {
	sql := "SELECT * from Link"
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

func (c Links) New() revel.Result {
	if c.Request.Method == http.MethodPost {
		form, err := c.Request.GetForm()
		if err != nil {
			revel.AppLog.Error("Form Error", err)
			return c.RenderError(err)
		}

		sql := "INSERT INTO Link (name, link) VALUES ($1, $2);"
		_, err = app.DB.Exec(sql, form.Get("name"), "http://example.com/x")
		if err != nil {
			revel.AppLog.Error("DB Error", err)
			return c.RenderError(err)
		}
		return c.Redirect(Links.List)
	}
	return c.Render()
}

func (c Links) Delete() revel.Result {
	if c.Request.Method == http.MethodPost {
		sql := "DELETE FROM Link where id = $1;"
		_, err := app.DB.Exec(sql, c.Params.Query.Get("id"))
		if err != nil {
			revel.AppLog.Error("DB Error", err)
			return c.RenderError(err)
		}
	}
	return c.Redirect(Links.List)
}
