package domain

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=4,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,max=1024"`
	CreatedAt time.Time
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for i, email := range emails {
		contacts[i].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}

	err := campaign.validate()

	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c Campaign) validate() error {

	v := validator.New()
	err := v.Struct(c)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			s := fmt.Sprintf("[%s] must have [%s] value [%s]", e.StructField(), e.Tag(), e.Param())
			return errors.New(s)
		}
	}

	return nil
}
