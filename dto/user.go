package dto

type Profile struct {
	Id       uint
	Email    string
	Username string
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Profile Profile
}

type SignupReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRes struct {
	User Profile
}
