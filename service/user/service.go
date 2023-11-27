package user

import "github.com/basliq/basliq-server-mvp/entity"

type Repository interface {
	IsEmailUnique(email string) (bool, error)
	CreateUser(user entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
