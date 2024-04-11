package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbUrl      string `json:"DbUrl"`
	Env        string `json:"Env"`
	ServerAddr string `json:"ServerAddr"`
}

func New() *Config {

	var cfg Config

	if err := cleanenv.ReadConfig("config.json", &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
