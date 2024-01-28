package campaing

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func TestNewService(t *testing.T) {

	command := CommandCampaign{
		Name:    "sell",
		Content: "content",
		Emails:  []string{"user@mail.com"},
	}

	repository := new(repositoryMock)
	service := NewService(repository)

	t.Run("should create a new service", func(t *testing.T) {

		repository.On("Save", mock.Anything).
			Times(1).
			Return(nil)

		err := service.Create(command)

		assert.Nil(t, err)
	})
}
