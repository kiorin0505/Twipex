package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Apikey string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("app.config")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Apikey: cfg.Section("trn").Key("api_key").String(),
	}
}
