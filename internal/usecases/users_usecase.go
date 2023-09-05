package usecases

import (
	"context"

	"github.com/maksewsha/wardrobe-back/internal/model"
	"github.com/maksewsha/wardrobe-back/internal/repository"
)

type UserUsecase struct {
	userRepo *repository.UserRepository
}

func (uc *UserUsecase) GetUserByLogin(ctx context.Context, login string) (*model.UserResponse, error) {
	resp, err := uc.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
