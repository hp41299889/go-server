package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Postgres struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

func EnvPostgres() Postgres {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var postgres Postgres
	postgres.Username = os.Getenv("POSTGRES_USERNAME")
	postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	postgres.Host = os.Getenv("POSTGRES_HOST")
	postgres.Port = os.Getenv("POSTGRES_PORT")
	postgres.Database = os.Getenv("POSTGRES_DATABASE")
	postgres.Charset = os.Getenv("POSTGRES_CHARSET")

	return postgres
}
