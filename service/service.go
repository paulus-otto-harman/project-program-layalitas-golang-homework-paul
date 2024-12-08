package service

import "homework/repository"

type Service struct {
}

func NewService(repo repository.Repository) Service {
	return Service{}
}
