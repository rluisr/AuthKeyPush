package models

import (
	"golang.org/x/oauth2"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Organization struct {
	Login string `json:"login"`
}

var g_h = githubConfig()

func githubConfig() *oauth2.Config {
	s := ReadSiteConfig()

	c := &oauth2.Config{
		ClientID:     s.GhClientID,
		ClientSecret: s.GhClientSec,
		RedirectURL:  s.GhRedirectUrl,
		Scopes:       []string{"read:org"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	return c
}

func GetGithubAuthUrl() (string) {
	return g_h.AuthCodeURL("")
}

// Check user organization and admin setting organization
// If its true, return true else return false
func GitHub(code string) (bool) {
	token, err := g_h.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}

	url := "https://api.github.com/user/orgs"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+token.AccessToken)

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	bytes := []byte(string(byteArray))

	g := []Organization{}
	json.Unmarshal(bytes, &g)

	if g[0].Login == ReadSiteConfig().GhAllowOrg {
		return true
	} else {
		return false
	}
}
