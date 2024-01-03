package user

import (
	"github.com/basliq/basliq-server-mvp/dto"
	"github.com/basliq/basliq-server-mvp/entity"
)

func (s Service) Signup(req dto.SignupReq) (dto.SignupRes, error) {
	createdUser, err := s.repo.CreateUser(entity.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return dto.SignupRes{}, err
	}

	return dto.SignupRes{User: createdUser}, nil
}
