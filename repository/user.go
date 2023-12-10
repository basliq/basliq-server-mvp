package repository

import "github.com/basliq/basliq-server-mvp/entity"

func (db DB) CreateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
