package campaing

type IService interface {
	Create(command CommandCampaign) error
}

type Service struct {
	repository IRepository
}

type CommandCampaign struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Emails  []string `json:"emails"`
}

func NewService(repository IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(command CommandCampaign) error {

	campaign, err := NewCampaign(command.Name, command.Content, command.Emails)

	if err != nil {
		return err
	}

	err = s.repository.Save(campaign)

	if err != nil {
		return err
	}

	return nil
}
