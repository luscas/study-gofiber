package models

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
