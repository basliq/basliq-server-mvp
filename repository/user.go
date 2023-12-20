package repository

import (
	"context"
	"fmt"
	"github.com/basliq/basliq-server-mvp/entity"
)

func (db DB) CreateUser(user entity.User) (entity.User, error) {
	userQueryRes := entity.User{}
	err := db.pool.QueryRow(context.Background(), `INSERT INTO users(email, username, password) VALUES($1, $2, $3);`, user.Email, user.Username, user.Password).Scan(&userQueryRes)

	if err != nil {
		fmt.Println(err)
		return entity.User{}, err
	}

	return userQueryRes, nil
}
