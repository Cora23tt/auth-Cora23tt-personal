package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	healthHandler "github.com/Cora23tt/auth-Cora23tt-personal/internal/rest/handlers/health"

	"github.com/Cora23tt/auth-Cora23tt-personal/internal/rest"
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
			return fmt.Errorf("failed to provide dependency: %w", err)
		}
	}

	if err := container.Invoke(func(server *rest.Server) {
		server.Init()
	}); err != nil {
		return fmt.Errorf("failed to invoke server: %w", err)
	}

	return fmt.Errorf("server failed to start: %w", container.Invoke(func(server *http.Server) error {
		return server.ListenAndServe()
	}))
}
