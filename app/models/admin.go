package models

import (
	"io/ioutil"
	"encoding/json"
	_"fmt"
)

func ReadSSHConfig() ([]Config) {
	buf, err := ioutil.ReadFile("./conf/ssh.json")
	if err != nil {
		panic(err)
	}
	bytes := []byte(string(buf))

	var c []Config
	json.Unmarshal(bytes, &c)

	return c
}