package main

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type config struct {
	RunAddress    string `env:"RUN_ADDRESS"  envDefault:"localhost:3200"`
	DataBaseURI   string `env:"DATABASE_URI" envDefault:"postgres://postgres:qaz@localhost:5432/mydb"`
	UserFileStore string `env:"USER_FILE_STORE"  envDefault:"file_store"`
}

func initConfig() *config {
	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		log.WithError(err).Fatal("init environment error")
	}
	return &cfg
}
