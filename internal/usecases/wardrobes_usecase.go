package usecases

import (
	"context"

	"github.com/maksewsha/wardrobe-back/internal/model"
	"github.com/maksewsha/wardrobe-back/internal/repository"
)

type wardrobeUseCase struct {
	repo repository.WardrobeRepository
}

func NewWardrobeUseCase(wr repository.WardrobeRepository) *wardrobeUseCase {
	return &wardrobeUseCase{
		repo: wr,
	}
}

func (w *wardrobeUseCase) CreateWardrobe(wd model.WardrobeDTO) (*uint64, error) {
	resp, err := w.repo.CreateNewWardrobe(context.Background(), wd)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wardrobeUseCase) GetAllWardrobesByUserId(id uint64) ([]model.WardrobeDAO, error) {
	resp, err := w.repo.GetAllWardrobes(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
