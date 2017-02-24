package controllers

import (
	"github.com/revel/revel"
	"AuthKeyPush/app/models"
)

type App struct {
	*revel.Controller
}

// Top page
func (c App) Index() revel.Result {
	site := models.ReadSiteConfig()
	title := site.Title

	// Check setting -> auth
	if site.AuthGithub == 1 && c.Session["login"] != "true" {
		// If session is false, show login page
		overview := site.AuthOverview
		url := models.GetGithubAuthUrl()
		msg := c.Session["msg"]
		login := 0

		return c.Render(title, overview, url, login, msg)
	}

	overview := site.TopOverview
	login := 1
	return c.Render(title, overview, login)
}

// Admin page
func (c App) Admin() revel.Result {
	site := models.ReadSiteConfig()
	ssh := models.ReadSSHConfig()

	title := site.Title
	overview := site.AdminOverview

	if site.AuthGithub == 1 && c.Session["login"] != "true" {
		return c.Redirect("/")
	}

	return c.Render(title, overview, ssh, site)
}

// Register new Authorized_keys
func (c App) Add(name, key string) revel.Result {
	site := models.ReadSiteConfig()

	if site.AuthGithub == 1 && c.Session["login"] != "true" {
		return c.Redirect("/")
	}

	models.ExecSSH(name, key)
	return c.Render(name, key)
}
