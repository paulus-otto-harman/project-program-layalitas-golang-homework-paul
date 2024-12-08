package service

import (
	"homework/domain"
	"homework/repository"
)

type UserService interface {
	Register(user *domain.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *domain.User) error {
	return s.repo.Create(user)
}
