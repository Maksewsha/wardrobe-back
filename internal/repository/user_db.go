package repository

import (
	"context"
	"fmt"

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
	user := &model.UserResponse{}
	err := userRepo.connection.QueryRow(ctx, "select login, email from users where login=$1", login).Scan(&user.Login, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserRepository) CreateNewUser(ctx context.Context, user model.UserAuthRequest) (*uint64, error) {
	var id uint64
	err := userRepo.connection.QueryRow(ctx, "insert into users (login, email, mobile_phone, password_hash) values ($1, $2, $3, $4) returning id", user.Login, user.Email, user.MobilePhone, user.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &id, nil
}
