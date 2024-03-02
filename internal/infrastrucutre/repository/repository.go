package repository

import (
	"github.com/vttrz/emailn/internal/domain"
)

type IRepository interface {
	Save(campaign *domain.Campaign) (*domain.Campaign, error)
}

type Repository struct {
	db any
}

func NewRepository(db any) IRepository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Save(campaign *domain.Campaign) (*domain.Campaign, error) {
	return nil, nil
}
