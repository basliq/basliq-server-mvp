package user

type ProfileReq struct {
	Id uint
}

type ProfileRes struct {
}

func (s Service) Profile(req ProfileReq) (ProfileRes, error) {
	return ProfileRes{}, nil
}
