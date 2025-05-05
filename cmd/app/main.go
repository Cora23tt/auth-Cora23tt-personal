package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	healthRepo "github.com/alif-academy-uz/auth-Cora23tt/internal/repository/health"
	healthHandler "github.com/alif-academy-uz/auth-Cora23tt/internal/rest/handlers/health"

	healthService "github.com/alif-academy-uz/auth-Cora23tt/internal/usecase/health"

	"github.com/alif-academy-uz/auth-Cora23tt/internal/rest"
)

func main() {
	var (
		port = "9999"
		host = "0.0.0.0"
	)

	if err := execute(host, port); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(host, port string) error {
	deps := []any{

		gin.New,

		healthRepo.NewRepo,
		healthService.NewService,
		healthHandler.NewHandler,

		rest.NewServer,

		func(server *rest.Server) *http.Server {
			return &http.Server{
				Addr:    net.JoinHostPort(host, port),
				Handler: server,
			}
		},
	}

	container := dig.New()
	for _, dep := range deps {
		if err := container.Provide(dep); err != nil {
			return err
		}
	}

	if err := container.Invoke(func(server *rest.Server) {
		server.Init()
	}); err != nil {
		return err
	}

	return container.Invoke(func(server *http.Server) error {
		return server.ListenAndServe()
	})
}
