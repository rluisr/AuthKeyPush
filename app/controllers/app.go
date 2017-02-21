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
	return c.Render()
}

// Admin page
func (c App) Admin() revel.Result {
	ssh := models.ReadSSHConfig()
	return c.Render(ssh)
}

// Register new Authorized_keys
func (c App) Add(name, key string) revel.Result {
	models.ExecSSH(name, key)
	return c.Render(name, key)
}
