package repository

import (
	"context"
	"fmt"
	"github.com/basliq/basliq-server-mvp/entity"
)

type CreatedUser struct {
	Id       uint
	Email    string
	Username string
}

func (db DB) CreateUser(user entity.User) (CreatedUser, error) {
	var userId uint = 0
	err := db.pool.QueryRow(context.Background(), `INSERT INTO users(email, username, password)
													   VALUES($1, $2, $3) RETURNING user_id;`, user.Email, user.Username, user.Password).Scan(&userId)

	if err != nil {
		fmt.Println(err)
		return CreatedUser{}, err
	}

	return CreatedUser{
		Id:       userId,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
