package repository

import (
	"github.com/vttrz/emailn/internal/domain"
)

type IRepository interface {
	Save(campaign *domain.Campaign) error
}

type Repository struct {
	db any
}

func NewRepository(db any) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Save(campaign *domain.Campaign) error {
	return nil
}
