package campaing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vttrz/emailn/internal/domain/campaign"
)

func TestNewCampaign(t *testing.T) {

	var (
		name    = "sell"
		content = "content"
		emails  = []string{"user@mail.com", "user2@mail.com"}
	)

	t.Run("should create a new campaign", func(t *testing.T) {

		c, _ := campaing.NewCampaign(name, content, emails)

		assert.NotEmpty(t, c)
		assert.NotNil(t, c.ID)
		assert.Equal(t, "sell", name)
		assert.Equal(t, "content", content)
		assert.Equal(t, len(emails), len(c.Contacts))
		assert.NotNil(t, c.CreatedAt)
	})

	t.Run("should not create a new campaign be null", func(t *testing.T) {
		c, _ := campaing.NewCampaign(name, content, emails)
		assert.NotNil(t, c.CreatedAt)
	})

	t.Run("create a campaign name must be required", func(t *testing.T) {
		_, err := campaing.NewCampaign("", content, emails)

		assert.NotNil(t, err)
		assert.Equal(t, "name is required", err.Error())
	})
}
