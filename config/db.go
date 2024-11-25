package database

import (
	"log"

	"github.com/trickqz/user-auth-token/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectBD() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Panic("Erro ao fazer a migração: ", err)
	}
}
