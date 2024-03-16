package domain_test

import (
	"github.com/vttrz/emailn/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {

	var (
		name    = "sell"
		content = "content"
		emails  = []string{"user@mail.com", "user2@mail.com"}
	)

	t.Run("should create a new campaign", func(t *testing.T) {

		c, _ := domain.NewCampaign(name, content, emails)

		assert.NotEmpty(t, c)
		assert.NotNil(t, c.ID)
		assert.Equal(t, "sell", name)
		assert.Equal(t, "content", content)
		assert.Equal(t, len(emails), len(c.Contacts))
		assert.NotNil(t, c.CreatedAt)
	})

	t.Run("should not create a new campaign be null", func(t *testing.T) {
		c, _ := domain.NewCampaign(name, content, emails)
		assert.NotNil(t, c.CreatedAt)
	})

	t.Run("campaign name must be greater than 4", func(t *testing.T) {
		_, err := domain.NewCampaign("", content, emails)

		assert.NotNil(t, err)
		assert.Equal(t, "[Name] must have [min] value [4]", err.Error())
	})

	t.Run("campaign content must be greater than 5", func(t *testing.T) {
		_, err := domain.NewCampaign(name, "", []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "[Content] must have [min] value [5]", err.Error())
	})

	t.Run("campaign contacts list must be greater than 1", func(t *testing.T) {
		_, err := domain.NewCampaign(name, content, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "[Contacts] must have [min] value [1]", err.Error())
	})

	t.Run("campaign should be start with 'pending' status", func(t *testing.T) {
		c, err := domain.NewCampaign(name, content, emails)

		assert.Empty(t, err)
		assert.Equal(t, domain.StatusPending, c.Status)
	})
}
