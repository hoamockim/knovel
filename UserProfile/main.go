package main

import (
	"knovel/userprofile/infrastructure"
	infra "knovel/userprofile/infrastructure/config"
	"knovel/userprofile/infrastructure/web"
	"knovel/userprofile/presentation"
	presentationConfig "knovel/userprofile/presentation/config"
	"knovel/userprofile/presentation/util/jwt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := infra.InitConfig(); err != nil {
		log.Fatal("Cannot read config: ", err)
	}
	presentationConfig.SetConfig(infra.GetConfig().TaskClientKey)
	jwt.InitJWT(infra.GetJwtKeys())
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
