package models

import (
	"log"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Key      string `json:"key"`
}

func ReadConfig() ([]Config) {
	buf, err := ioutil.ReadFile("./conf/ssh.json")
	if err != nil {
		panic(err)
	}
	bytes := []byte(string(buf))

	var c []Config
	json.Unmarshal(bytes, &c)

	return c
}

func ExecSSH(posted_name, posted_key string) {
	c := ReadConfig()

	for _, c := range c {
		if c.Key == "" { // Password Authentication
			config := &ssh.ClientConfig{
				User: c.User,
				Auth: []ssh.AuthMethod{
					ssh.Password(c.Password),
				},
			}

			conn, err := ssh.Dial("tcp", c.Host+":"+c.Port, config)
			if err != nil {
				panic(err)
			}
			defer conn.Close()

			session, err := conn.NewSession()
			if err != nil {
				log.Println(err)
			}
			defer session.Close()

			session.Run("echo -e \"\n#\n# " + posted_name + "\n#\n" + posted_key + "\" >> .ssh/authorized_keys")

		} else { // Public Key Authentication
			buf, err := ioutil.ReadFile(c.Key)
			if err != nil {
				panic(err)
			}

			key, err := ssh.ParsePrivateKey(buf)
			if err != nil {
				panic(err)
			}

			config := &ssh.ClientConfig{
				User: c.User,
				Auth: []ssh.AuthMethod{
					ssh.PublicKeys(key),
				},
			}

			conn, err := ssh.Dial("tcp", c.Host+":"+c.Port, config)
			if err != nil {
				panic(err)
			}
			defer conn.Close()

			session, err := conn.NewSession()
			if err != nil {
				log.Println(err)
			}
			defer session.Close()

			session.Run("echo -e \"\n#\n# " + posted_name + "\n#\n" + posted_key + "\" >> .ssh/authorized_keys")
		}
	}
}
