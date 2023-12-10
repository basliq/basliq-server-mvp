package user

import "github.com/basliq/basliq-server-mvp/entity"

type RegisterReq struct {
	Username string
	Email    string
	Password string
}

type RegisterRes struct {
	User entity.User
}

func (s Service) Register(req RegisterReq) (RegisterRes, error) {
	createdUser, err := s.repo.CreateUser(entity.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return RegisterRes{}, err
	}

	return RegisterRes{User: createdUser}, nil
}
