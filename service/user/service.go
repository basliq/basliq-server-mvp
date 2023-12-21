package user

import (
	"github.com/basliq/basliq-server-mvp/entity"
	"github.com/basliq/basliq-server-mvp/repository"
)

type Repository interface {
	//IsUsernameUnique(username string) (bool, error)
	//IsEmailUnique(email string) (bool, error)
	//IsUsernameAndEmailUnique(email, username string) (bool, error)
	CreateUser(user entity.User) (repository.CreatedUser, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
