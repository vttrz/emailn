package campaing

type IRepository interface {
	Save(campaign *Campaign) error
}

type Repository struct {
	db any
}

func NewRepository(db any) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Save(campaign *Campaign) error {
	return nil
}
