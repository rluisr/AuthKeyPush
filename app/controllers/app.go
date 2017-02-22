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
	s := models.ReadSiteConfig()
	title := s.Title
	overview := s.TopOverview

	return c.Render(title, overview)
}

// Admin page
func (c App) Admin() revel.Result {
	site := models.ReadSiteConfig()
	ssh := models.ReadSSHConfig()

	title := site.Title
	overview := site.AdminOverview

	return c.Render(title, overview, ssh, site)
}

// Register new Authorized_keys
func (c App) Add(name, key string) revel.Result {
	models.ExecSSH(name, key)
	return c.Render(name, key)
}
