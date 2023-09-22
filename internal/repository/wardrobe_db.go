package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/maksewsha/wardrobe-back/internal/model"
)

type WardrobeRepository struct {
	connection *pgx.Conn
}

func NewWardrobeRepository(conn *pgx.Conn) *WardrobeRepository {
	return &WardrobeRepository{
		connection: conn,
	}
}

func (wr *WardrobeRepository) CreateNewWardrobe(ctx context.Context, wardrobe model.WardrobeDTO) (*uint64, error) {
	var id uint64
	err := wr.connection.QueryRow(ctx, "insert into wardrobes (title, description, poster) values $1, $2, $3 returning id", wardrobe.Title, wardrobe.Description, wardrobe.Poster).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (wr *WardrobeRepository) GetAllWardrobes(ctx context.Context, id uint64) ([]model.WardrobeDAO, error) {
	rows, err := wr.connection.Query(ctx, "select * from wardrobes join wardrobes_users on wardrobes.id = wardrobes_users.wardrobe_id where wardrobes_users.user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []model.WardrobeDAO
	for rows.Next() {
		var wardrobe model.WardrobeDAO
		err := rows.Scan(&wardrobe.Id, &wardrobe.Title, &wardrobe.Description, &wardrobe.Poster, &wardrobe.CreatedTime)
		if err != nil {
			return nil, err
		}
		data = append(data, wardrobe)
	}
	return data, nil
}
