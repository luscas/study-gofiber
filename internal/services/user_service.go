package services

type UserServiceInterface interface {
	GetName() string
	GetEmail() string
}

type UserService struct {
	Name     string
	Email    string
	Password string
}

func (u *UserService) GetName() string {
	return u.Name
}

func (u *UserService) GetEmail() string {
	return u.Email
}
