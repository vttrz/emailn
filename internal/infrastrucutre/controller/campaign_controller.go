package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vttrz/emailn/foundation"
	"github.com/vttrz/emailn/internal/infrastrucutre/usecase"
	"net/http"
)

type CampaignController struct {
	createCampaign usecase.CreateCampaignUseCase
}

func NewCampaignController(createCampaign usecase.CreateCampaignUseCase) *CampaignController {
	return &CampaignController{
		createCampaign: createCampaign,
	}
}

func (c CampaignController) Create(ctx *gin.Context) {

	var cmd usecase.CreateCampaignCommand

	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		foundation.AbortWithStatusError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := c.createCampaign.Execute(cmd)

	if err != nil {
		foundation.AbortWithStatusError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	foundation.Success(ctx, http.StatusCreated, response)
	return
}

func (c CampaignController) List(ctx *gin.Context) {
	foundation.Success(ctx, http.StatusOK, nil)
}
