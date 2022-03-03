package main

import (
	"github.com/rpmcdougall/es-demo/internal/pkg/config"
	"github.com/rpmcdougall/es-demo/internal/pkg/es"
	"github.com/rpmcdougall/es-demo/internal/pkg/handlers"
	"github.com/rpmcdougall/es-demo/internal/pkg/logging"
	"github.com/rpmcdougall/es-demo/internal/pkg/repository"
	"github.com/rpmcdougall/es-demo/internal/pkg/serve"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		repository.Module(),
		handlers.Module(),
		es.Module(),
		config.Module(),
		logging.Module(),
		fx.Invoke(serve.NewServer),
	)

	app.Run()
}
