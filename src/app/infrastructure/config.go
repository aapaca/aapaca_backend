package infrastructure

import "github.com/kelseyhightower/envconfig"

type DBConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

func NewDBConfig() DBConfig {
	var c DBConfig
	err := envconfig.Process("mysql", &c)
	if err != nil {
		panic(err)
	}
	return c
}
