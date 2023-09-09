package model

type UserAuthRequest struct {
	Login       string `json:"login"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobile_phone"`
	Password    string `json:"password"`
}

type UserResponse struct {
	Login string `json:"login"`
	Email string `json:"email"`
}
