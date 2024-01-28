package campaing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vttrz/emailn/internal/domain/campaign"
)

func TestNewCampaign(t *testing.T) {

	t.Run("should create a new campaign", func(t *testing.T) {
		name := "sell"
		content := "content"
		emails := []string{"user@mail.com", "user2@mail.com"}

		c := campaing.NewCampaign(name, content, emails)

		assert.NotEmpty(t, c)
		assert.NotNil(t, c.ID)
		assert.Equal(t, "sell", name)
		assert.Equal(t, "content", content)
		assert.Equal(t, len(emails), len(c.Contacts))
		assert.NotNil(t, c.CreatedAt)

	})

}
