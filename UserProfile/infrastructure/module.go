package infrastructure

import (
	"knovel/userprofile/application"
	"knovel/userprofile/domain/repositories"
	"knovel/userprofile/infrastructure/db"
	"knovel/userprofile/presentation"
	"knovel/userprofile/presentation/handler"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(db.NewDBContext)
	return container
}

func ProvisonDependencies(container *dig.Container) *dig.Container {
	container.Provide(repositories.NewRbacRepository)
	container.Provide(repositories.NewUserProfileRepository)
	container.Provide(application.NewApplication)
	container.Provide(handler.NewAuthHandler)
	container.Provide(handler.NewSysHandler)
	container.Provide(presentation.InitRouter)
	return container
}
