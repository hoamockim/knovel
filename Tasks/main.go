package main

import (
	"knovel/tasks/infrastructure"
	infraConfig "knovel/tasks/infrastructure/config"
	"knovel/tasks/infrastructure/web"
	"knovel/tasks/presentation"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := infraConfig.InitConfig(); err != nil {
		log.Fatal("Cannot read config: ", err)
	}
	if err := runApp(); err != nil {
		log.Fatal(err)
	}
}

func runApp() error {
	c := infrastructure.BuildContainer()
	c = infrastructure.ProvisonDependencies(c)
	return c.Invoke(func(router presentation.Router) error {
		return web.Run(router)
	})
}
