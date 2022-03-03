package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"log"
	"os"
)

type Config struct {
	EsUsername      string   `koanf:"es_user"`
	EsPassword      string   `koanf:"es_pass"`
	LogLevel        string   `koanf:"log_level"`
	LogCallerEnable bool     `koanf:"log_caller_enable"`
	ElasticHosts    []string `koanf:"elastic_hosts"`
	EsCert          []byte   `koanf:"es_cert"`
}

var (
	k      = koanf.New(".")
	parser = yaml.Parser()
)

func NewConfig() (*Config, error) {

	if err := k.Load(file.Provider("settings.yaml"), parser); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	var bindConfig Config
	err := k.Unmarshal("", &bindConfig)
	if err != nil {
		return nil, err
	}

	certData, err := os.ReadFile("ca.crt")
	if err != nil {
		return nil, err
	}

	return &Config{
		EsUsername:      bindConfig.EsUsername,
		EsPassword:      bindConfig.EsPassword,
		LogLevel:        bindConfig.LogLevel,
		LogCallerEnable: bindConfig.LogCallerEnable,
		ElasticHosts:    bindConfig.ElasticHosts,
		EsCert:          certData}, nil
}
