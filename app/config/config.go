package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Database struct {
		Host     string
		Name     string
		Username string
		Password string
	}
	Admin struct {
		Name     string
		Email    string
		Password string
		Profile  string
	}
	KeyJWT struct {
		Author string
		Admin  string
	}
}

func ReadJsonConfig() Config {
	content, err := ioutil.ReadFile("./app/config/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload Config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}
