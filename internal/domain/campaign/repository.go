package campaing

type Repository interface {
	Save(campaign *Campaign) error
}
