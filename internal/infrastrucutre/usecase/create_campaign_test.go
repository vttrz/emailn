package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vttrz/emailn/internal/domain"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *domain.Campaign) (*domain.Campaign, error) {
	args := r.Called(campaign)
	return nil, args.Error(0)
}

func TestNewCreateCampaignUseCase(t *testing.T) {

	command := CreateCampaignCommand{
		Name:    "sell",
		Content: "content",
		Emails:  []string{"user@mail.com"},
	}

	repository := new(repositoryMock)
	uc := NewCreateCampaignUseCase(repository)

	t.Run("should create a new service", func(t *testing.T) {

		repository.On("Save", mock.Anything).
			Times(1).
			Return(nil)

		_, err := uc.Execute(command)

		assert.Nil(t, err)
	})
}
