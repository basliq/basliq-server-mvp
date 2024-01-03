package user

import "github.com/basliq/basliq-server-mvp/dto"

func (s Service) Login(req dto.LoginReq) (dto.LoginRes, error) {
	user, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		return dto.LoginRes{}, err
	}

	return dto.LoginRes{Profile: user}, nil
}
