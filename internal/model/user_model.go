package model

type UserAuthRequest struct {
	Login       string
	Email       string
	MobilePhone string
	Password    string
}

type UserResponse struct {
	Login string
	Email string
}
