package campaing

type Service struct {
	repository Repository
}

type CommandCampaign struct {
	Name    string
	Content string
	Emails  []string
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(command CommandCampaign) error {

	_, err := NewCampaign(command.Name, command.Content, command.Emails)

	if err != nil {
		return err
	}

	return nil
}
