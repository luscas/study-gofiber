package database

import (
	"fmt"
	"os"

	"api.droppy.com.br/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("variável de ambiente DATABASE_URL não definida")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
