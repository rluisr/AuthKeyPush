package controllers

import (
	"github.com/revel/revel"
	"AuthKeyPush/app/models"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Github(code string) revel.Result {
	login := models.GitHub(code)

	if login == true {
		c.Session["login"] = "true"
	} else {
		c.Session["login"] = "false"
		c.Session["msg"] = "Login faied. Check conf/site.json or README.md"
	}

	return c.Redirect("/")
}
