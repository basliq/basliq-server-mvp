package user

import (
	"github.com/basliq/basliq-server-mvp/dto"
	"github.com/basliq/basliq-server-mvp/entity"
)

type Repository interface {
	CreateUser(user entity.User) (dto.Profile, error)
	FindUserByEmail(email string) (dto.Profile, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
