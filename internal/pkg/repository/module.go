package repository

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Provide(NewLiveRepository)
}
