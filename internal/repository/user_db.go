package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/maksewsha/wardrobe-back/internal/model"
)

type UserRepository struct {
	connection *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		connection: conn,
	}
}

func (userRepo *UserRepository) GetUserByLogin(ctx context.Context, login string) (*model.UserResponse, error) {
	var user *model.UserResponse
	err := userRepo.connection.QueryRow(ctx, "select login, email from users where login=$1", login).Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
