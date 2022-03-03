package serve

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rpmcdougall/es-demo/internal/pkg/handlers"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type NewServerParams struct {
	fx.In

	Handlers  *handlers.ApiHandlers
	Lifecycle fx.Lifecycle
	Logger    *logrus.Logger
}

func NewServer(p NewServerParams) *http.Server {
	r := gin.Default()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	v1Group := r.Group("/v1")
	for _, handler := range p.Handlers.Handlers {
		v1Group.Handle(handler.Method, handler.Path, handler.HandlerFunc)
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			p.Logger.Print("Starting HTTP server.")
			go func() {
				// service connections
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Logger.Print("Stopping HTTP server.")
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
