package repository

import (
	"context"
	"github.com/basliq/basliq-server-mvp/dto"
	"github.com/basliq/basliq-server-mvp/entity"
)

func (db DB) CreateUser(user entity.User) (dto.Profile, error) {
	var userId uint = 0
	err := db.pool.QueryRow(context.Background(), `INSERT INTO users(email, username, password)
													   VALUES($1, $2, $3) RETURNING user_id;`, user.Email, user.Username, user.Password).Scan(&userId)

	if err != nil {
		return dto.Profile{}, err
	}

	return dto.Profile{
		Id:       userId,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (db DB) FindUserByEmail(email string) (dto.Profile, error) {
	var username string
	var password string
	err := db.pool.QueryRow(context.Background(), `SELECT username, password FROM users WHERE email = $1`, email).Scan(&username, &password)
	if err != nil {
		return dto.Profile{}, err
	}

	// TODO - find a better structure for dto, entity and other types
	return dto.Profile{
		Id:       0,
		Email:    email,
		Username: username,
	}, nil
}
