package infrastructure

import (
	"knovel/tasks/application"
	"knovel/tasks/domain/repositories"
	infra "knovel/tasks/infrastructure/client"
	"knovel/tasks/infrastructure/db"
	"knovel/tasks/presentation"
	"knovel/tasks/presentation/handler"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(db.NewDBContext)
	return container
}

func ProvisonDependencies(container *dig.Container) *dig.Container {
	container.Provide(repositories.NewTaskRepository)
	container.Provide(application.NewApplication)
	container.Provide(infra.NewAuthorClient)
	container.Provide(handler.NewTaskHandler)
	container.Provide(presentation.InitRouter)

	return container
}
