package models

import (
	"io/ioutil"
	"encoding/json"
)

type Site struct {
	AuthGithub    int `json:"auth-github-login"`
	GhClientID    string `json:"github-client-id"`
	GhClientSec   string `json:"github-client-secret"`
	GhRedirectUrl string `json:"github-redirect-url"`
	GhAllowOrg    string `json:"github-allow-organization"`
	Title         string `json:"title"`
	TopOverview   string `json:"top-overview"`
	AuthOverview  string `json:"auth-overview"`
	AdminOverview string `json:"admin-overview"`
}

func ReadSiteConfig() (Site) {
	buf, err := ioutil.ReadFile("./conf/site.json")
	if err != nil {
		panic(err)
	}
	bytes := []byte(buf)

	var s Site
	json.Unmarshal(bytes, &s)

	return s
}
