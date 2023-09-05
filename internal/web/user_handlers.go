package web

import "github.com/maksewsha/wardrobe-back/internal/usecases"

type UserHandlers struct {
	userUseCase *usecases.UserUsecase
}

func NewUserHandlers(uc *usecases.UserUsecase) *UserHandlers {
	return &UserHandlers{
		userUseCase: uc,
	}
}
