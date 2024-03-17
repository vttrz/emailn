package usecase

import (
	"github.com/vttrz/emailn/internal/infrastrucutre/repository"
)

type (
	GetCampaignByIDCommand struct {
		ID string
	}

	GetByIDCampaignResponse struct {
		ID      string
		Name    string
		Content string
		Status  string
	}

	GetByIDCampaignUseCase interface {
		Execute(command GetCampaignByIDCommand) (*GetByIDCampaignResponse, error)
	}

	getByIDCampaignUseCase struct {
		repository repository.IRepository
	}
)

func NewGetCampaignByIDUseCase(repository repository.IRepository) GetByIDCampaignUseCase {
	return &getByIDCampaignUseCase{
		repository: repository,
	}
}

func (c getByIDCampaignUseCase) Execute(command GetCampaignByIDCommand) (*GetByIDCampaignResponse, error) {

	campaign, err := c.repository.GetByID(command.ID)

	if err != nil {
		return nil, err
	}

	return &GetByIDCampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
