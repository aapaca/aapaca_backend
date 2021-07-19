package infrastructure

import (
	"github.com/kelseyhightower/envconfig"
)

type MysqlConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

func NewMysqlConfig() MysqlConfig {
	var c MysqlConfig
	err := envconfig.Process("mysql", &c)
	if err != nil {
		panic(err)
	}
	return c
}

type CloudSqlConfig struct {
	User           string
	Password       string
	Database       string
	ConnectionName string
}

func NewCloudSqlConfig() CloudSqlConfig {
	var c CloudSqlConfig
	err := envconfig.Process("db", &c)
	if err != nil {
		panic(err)
	}
	return c
}
