package services

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-server/configs"
	"go-server/models"
)

var Db *gorm.DB

type User models.User

func PostgresConnect() {
	envPostgres := configs.EnvPostgres()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		envPostgres.Host, envPostgres.Username, envPostgres.Password,
		envPostgres.Database, envPostgres.Port)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Connect to postgres failed:", err)
	}

	Db.AutoMigrate(&User{})
	log.Println("Connect to postgres successfully")
	log.Println(Db)
}
