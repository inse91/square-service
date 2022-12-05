package config

import (
	c "github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

type config struct {
	Port int `yaml:"port"`
}

var cfg *config = &config{}

func GetConfig() (*config, error) {

	conf := c.New()
	yaml := feeder.Yaml{Path: "internal/config/config.yaml"}

	conf.AddFeeder(yaml)
	conf.AddStruct(cfg)
	err := conf.Feed()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
