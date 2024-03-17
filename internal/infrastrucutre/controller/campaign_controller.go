package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vttrz/emailn/foundation"
	"github.com/vttrz/emailn/internal/infrastrucutre/usecase"
	"net/http"
)

type CampaignController struct {
	createCampaign  usecase.CreateCampaignUseCase
	getCampaignByID usecase.GetByIDCampaignUseCase
}

func NewCampaignController(createCampaign usecase.CreateCampaignUseCase, getCampaignByID usecase.GetByIDCampaignUseCase) *CampaignController {
	return &CampaignController{
		createCampaign:  createCampaign,
		getCampaignByID: getCampaignByID,
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

func (c CampaignController) GetByID(ctx *gin.Context) {

	var cmd usecase.GetCampaignByIDCommand

	cmd.ID = ctx.Param("id")

	response, err := c.getCampaignByID.Execute(cmd)

	if err != nil {
		foundation.AbortWithStatusError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	foundation.Success(ctx, http.StatusOK, response)
	return
}
