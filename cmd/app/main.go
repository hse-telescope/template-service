package main

import (
	"os"

	"github.com/hse-telesope/template-service/internal/config"
	"github.com/hse-telesope/template-service/internal/providers/people"
	"github.com/hse-telesope/template-service/internal/repository/facade"
	"github.com/hse-telesope/template-service/internal/repository/storage"
	"github.com/hse-telesope/template-service/internal/server"
)

func main() {
	configPath := os.Args[1]
	conf, err := config.Parse(configPath)
	if err != nil {
		panic(err)
	}

	storage, err := storage.New(conf.DB.GetDBURL(), conf.DB.MigrationsPath)
	if err != nil {
		panic(err)
	}

	facade := facade.New(storage)

	provider := people.New(facade)

	s := server.New(conf, provider)
	panic(s.Start())
}
