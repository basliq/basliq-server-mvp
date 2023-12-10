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
	return RegisterRes{}, nil
}
