package internal

import (
	"github.com/vttrz/emailn/internal/infrastrucutre/controller"
	"github.com/vttrz/emailn/internal/infrastrucutre/repository"
	"github.com/vttrz/emailn/internal/infrastrucutre/usecase"
)

type Application struct {
	CampaignController *controller.CampaignController
}

func NewApplication() *Application {

	repo := repository.NewRepository(nil)

	return &Application{
		CampaignController: controller.NewCampaignController(
			usecase.NewCreateCampaignUseCase(repo),
		),
	}
}
