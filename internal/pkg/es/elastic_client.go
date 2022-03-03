package es

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rpmcdougall/es-demo/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type ElasticClient struct {
	API    *elasticsearch.Client
	config *config.Config
	logger *logrus.Logger
}

type ElasticClientParams struct {
	fx.In

	Config *config.Config
	Logger *logrus.Logger
}

func NewLiveElasticClient(p ElasticClientParams) (*ElasticClient, error) {

	cfg := elasticsearch.Config{
		Username:  p.Config.EsUsername,
		Password:  p.Config.EsPassword,
		Addresses: p.Config.ElasticHosts,
		CACert:    p.Config.EsCert,
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticClient{
		API:    es,
		config: p.Config,
		logger: p.Logger,
	}, nil
}
