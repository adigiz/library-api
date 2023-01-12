package repository

import (
	"errors"
	"fmt"
	"library-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers(query string) ([]*models.User, error)
	MatchingCredential(email, password string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type UROpts struct {
	DB *gorm.DB
}

func NewUserRepository(opts *UROpts) UserRepository {
	return &userRepository{db: opts.DB}
}

func (u *userRepository) FindUsers(name string) ([]*models.User, error) {
	var users []*models.User
	result := u.db.Where("Name ILIKE ?", fmt.Sprintf("%%%v%%", name)).Find(&users)
	return users, result.Error
}

func (u *userRepository) MatchingCredential(email, password string) (*models.User, error) {
	var user *models.User
	query := &models.User{
		Email:    email,
		Password: password,
	}
	err := u.db.Where(query).First(&user).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, err
	}
	return user, err
}
