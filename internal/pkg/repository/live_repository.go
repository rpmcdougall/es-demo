package repository

import (
	"encoding/json"
	"github.com/rpmcdougall/es-demo/internal/pkg/es"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"log"
)

type LiveRepository struct {
	elasticClient *es.ElasticClient
	logger        *logrus.Logger
}

func (l *LiveRepository) GetInfo() (map[string]interface{}, error) {

	var r map[string]interface{}

	info, err := l.elasticClient.API.Info()
	if err != nil {
		return nil, err
	}

	defer info.Body.Close()

	if err := json.NewDecoder(info.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	return r, nil
}

type LiveRepositoryParams struct {
	fx.In

	ElasticClient *es.ElasticClient
	Logger        *logrus.Logger
}

func NewLiveRepository(p LiveRepositoryParams) *LiveRepository {
	return &LiveRepository{
		elasticClient: p.ElasticClient,
		logger:        p.Logger,
	}
}
