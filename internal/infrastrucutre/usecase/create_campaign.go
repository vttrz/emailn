package usecase

import (
	"github.com/vttrz/emailn/internal/domain"
	"github.com/vttrz/emailn/internal/infrastrucutre/repository"
)

type (
	CreateCampaignCommand struct {
		Name    string   `json:"name"`
		Content string   `json:"content"`
		Emails  []string `json:"emails"`
	}

	CreateCampaignResponse struct {
		Name    string   `json:"name"`
		Content string   `json:"content"`
		Emails  []string `json:"emails"`
	}

	CreateCampaignUseCase interface {
		Execute(command CreateCampaignCommand) (*CreateCampaignResponse, error)
	}

	createCampaignUseCase struct {
		repository repository.IRepository
	}
)

func NewCreateCampaignUseCase(repository repository.IRepository) CreateCampaignUseCase {
	return &createCampaignUseCase{
		repository: repository,
	}
}

func (c createCampaignUseCase) Execute(command CreateCampaignCommand) (*CreateCampaignResponse, error) {

	campaign, _ := domain.NewCampaign(command.Name, command.Content, command.Emails)

	_, err := c.repository.Save(campaign)

	if err != nil {
		return nil, err
	}

	return &CreateCampaignResponse{}, nil
}
