package config

import (
	"github.com/kelseyhightower/envconfig"
	"os"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetDBEnv() (DBEnv, error) {
	var dbEnv DBEnv
	err := envconfig.Process("DATABASE", &dbEnv)
	return dbEnv, err
}
