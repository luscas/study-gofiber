package interfaces

import "api.droppy.com.br/internal/models"

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
}
