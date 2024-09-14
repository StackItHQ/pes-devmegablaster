package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfig struct {
	DbConfig
	ApiConfig
}

type DbConfig struct {
	HOST   string
	USER   string
	PASS   string
	DBNAME string
	PORT   string
}

type ApiConfig struct {
	PORT string
}

var Cfg *ServerConfig

func Init() {
	Cfg = &ServerConfig{
		DbConfig: DbConfig{
			HOST:   os.Getenv("POSTGRES_HOST"),
			USER:   os.Getenv("POSTGRES_USER"),
			PASS:   os.Getenv("POSTGRES_PASS"),
			DBNAME: os.Getenv("POSTGRES_DBNAME"),
			PORT:   os.Getenv("POSTGRES_PORT"),
		},
		ApiConfig: ApiConfig{
			PORT: os.Getenv("API_PORT"),
		},
	}
}
