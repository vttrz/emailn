package campaing

import (
	"github.com/vttrz/emailn/internal/domain"
	repository2 "github.com/vttrz/emailn/internal/infrastrucutre/repository"
)

type IService interface {
	Create(command CommandCampaign) error
}

type Service struct {
	repository repository2.IRepository
}

type CommandCampaign struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Emails  []string `json:"emails"`
}

func NewService(repository repository2.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(command CommandCampaign) error {

	campaign, err := domain.NewCampaign(command.Name, command.Content, command.Emails)

	if err != nil {
		return err
	}

	err = s.repository.Save(campaign)

	if err != nil {
		return err
	}

	return nil
}
