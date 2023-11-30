package user

type LoginReq struct {
	Username string
	Password string
}

type LoginRes struct {
}

func (s Service) Login(req LoginReq) (LoginRes, error) {
	return LoginRes{}, nil
}
