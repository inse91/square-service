package config

import (
	c "github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

type MongoDB struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Login      string `yaml:"login"`
	Password   string `yaml:"password"`
	DataBase   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

type PostgresDB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DataBase string `yaml:"database"`
}

type config struct {
	Port       string     `yaml:"port"`
	MongoDB    MongoDB    `yaml:"mongo"`
	PostgresDB PostgresDB `yaml:"postgres"`
}

var cfg *config = &config{}

func GetConfig() (*config, error) {

	conf := c.New()

	conf.
		AddFeeder(feeder.Yaml{Path: "config.yaml"}).
		AddStruct(cfg)

	if err := conf.Feed(); err != nil {
		return nil, err
	}

	return cfg, nil

}
