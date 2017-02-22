package models

import (
	"io/ioutil"
	"encoding/json"
)

type Site struct {
	Title         string `json:"title"`
	TopOverview   string `json:"top-overview"`
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
