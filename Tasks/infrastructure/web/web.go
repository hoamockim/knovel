package web

import (
	"fmt"
	"knovel/tasks/infrastructure/config"
	"knovel/tasks/presentation"
	"net/http"
)

func Run(router presentation.Router) error {
	server := &http.Server{
		Addr:           config.GetServerAddress(),
		Handler:        router.GetRouter(),
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Service running on port: ", config.GetServerAddress())
	return server.ListenAndServe()
}
