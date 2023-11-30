package user

import "github.com/basliq/basliq-server-mvp/entity"

type Repository interface {
	IsEmailUnique(email string) (bool, error)
	CreateUser(user entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, bool, error)
	GetUserById(id uint)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
