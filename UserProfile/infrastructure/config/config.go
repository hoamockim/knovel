package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type AppConfig struct {
	RunMode       string `default:"DEBUG"`
	Port          int    `required:"true" envconfig:"PORT" default:"8081"`
	Db            Postgres
	Jwt           JwtKey
	TaskClientKey string `required:"true" envconfig:"TASK_CLIENT_KEY"`
}

type Postgres struct {
	Host         string `required:"true" envconfig:"PG_HOST"`
	Port         int    `required:"true" envconfig:"PG_PORT"`
	User         string `required:"true" envconfig:"PG_USER"`
	Password     string `required:"true" envconfig:"PG_PASSWORD"`
	DB           string `required:"true" envconfig:"PG_DB"`
	SearchPath   string `default:"public"`
	SSLMode      string `default:"disable"`
	MaxOpenConn  int    `required:"true" default:"25" envconfig:"PG_MAX_IDLE_CONN"`
	MaxIdleConn  int    `required:"true" default:"25" envconfig:"PG_MAX_OPEN_CONN"`
	ConnLifeTime int    `required:"true" default:"5"  envconfig:"PG_CONN_LIFE_TIME"` // time in minutes

}

var appConfig AppConfig

func GetConfig() AppConfig {
	return appConfig
}
func GetServerAddress() string {
	return fmt.Sprintf(":%d", appConfig.Port)
}

func GetJwtKeys() (string, string) {
	return appConfig.Jwt.PubPath, appConfig.Jwt.PriPath
}

func InitConfig() (err error) {

	if err = dotenv.Load(); err != nil {
		return
	}
	err = envconfig.Process("", &appConfig)

	return
}
