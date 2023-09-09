package usecases

import (
	"context"

	"github.com/maksewsha/wardrobe-back/internal/model"
)

type (
	User interface {
		GetUserByLogin(context.Context, string) (*model.UserResponse, error)
		CreateNewUser(context.Context, model.UserAuthRequest) (*uint64, error)
	}
)
