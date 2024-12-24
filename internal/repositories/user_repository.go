package repositories

import (
	"api.droppy.com.br/internal/models"
	"api.droppy.com.br/internal/services"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT id, name, email, created_at, updated_at FROM users").Scan(&users).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?", id).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user services.UserService) error {
	err := r.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password).Error
	if err != nil {
		return err
	}
	return nil
}
