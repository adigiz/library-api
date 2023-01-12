package services

import (
	"library-api/models"
	repository "library-api/repositories"
)

type UserService interface {
	GetUsers(name string) ([]*models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

type USOpts struct {
	UserRepository repository.UserRepository
}

func NewUserService(opts *USOpts) UserService {
	return &userService{userRepository: opts.UserRepository}
}

func (u *userService) GetUsers(name string) ([]*models.User, error) {
	return u.userRepository.FindUsers(name)
}
