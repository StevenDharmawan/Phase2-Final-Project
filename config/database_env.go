package config

type DBEnv struct {
	Host     string `envconfig:"HOST"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Port     string `envconfig:"PORT"`
	DBName   string `envconfig:"DBNAME"`
}
