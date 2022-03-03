package logging

import (
	"github.com/rpmcdougall/es-demo/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type params struct {
	fx.In

	Config *config.Config
}

func NewLogger(p params) (*logrus.Logger, error) {

	newLogger := logrus.New()

	level, err := logrus.ParseLevel(p.Config.LogLevel)
	if err != nil {
		return nil, err
	}
	newLogger.SetLevel(level)

	newLogger.SetReportCaller(p.Config.LogCallerEnable)

	return newLogger, nil
}
