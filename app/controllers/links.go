package controllers

import (
	"exercicio-golang-03/app"
	"net/http"

	"github.com/didi/gendry/scanner"
	"github.com/mniak/superlink-sdk"
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
		var price float64
		var name string
		var linktype string
		var shippingtype string
		c.Params.Bind(&price, "price")
		c.Params.Bind(&name, "name")
		c.Params.Bind(&linktype, "type")
		c.Params.Bind(&shippingtype, "shipping.type")

		link, err := app.SuperLink.CreateLink(superlink.Link{
			Name:  name,
			Price: int(price * 100),
			Type:  linktype,
			Shipping: superlink.Shipping{
				Type: shippingtype,
			},
		})
		if err != nil {
			revel.AppLog.Error("Form Error", "Error", err)
			return c.RenderError(err)
		}

		sql := "INSERT INTO Link (name, link) VALUES ($1, $2);"
		_, err = app.DB.Exec(sql, link.Name, link.ShortURL)
		if err != nil {
			revel.AppLog.Error("DB Error", "Error", err)
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
			revel.AppLog.Error("DB Error", "Error", err)
			return c.RenderError(err)
		}
	}
	return c.Redirect(Links.List)
}
