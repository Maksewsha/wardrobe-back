package usecases

import (
	"context"

	"github.com/maksewsha/wardrobe-back/internal/model"
	"github.com/maksewsha/wardrobe-back/internal/repository"
)

type userUseCase struct {
	userRepo *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *userUseCase {
	return &userUseCase{userRepo: repo}
}

func (uc *userUseCase) GetUserByLogin(ctx context.Context, login string) (*model.UserResponse, error) {
	resp, err := uc.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *userUseCase) CreateNewUser(ctx context.Context, user model.UserAuthRequest) (*uint64, error) {
	resp, err := uc.userRepo.CreateNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
