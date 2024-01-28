package campaing

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedAt time.Time
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))

	for i, email := range emails {
		contacts[i].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}
}
