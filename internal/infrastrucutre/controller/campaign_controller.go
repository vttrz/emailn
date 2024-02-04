package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vttrz/emailn/foundation"
	campaing "github.com/vttrz/emailn/internal/domain/campaign"
	"net/http"
)

type Handlers struct {
	service    campaing.IService
	repository campaing.IRepository
}

func NewHandlers(service campaing.IService, repository campaing.IRepository) *Handlers {
	return &Handlers{
		service:    service,
		repository: repository,
	}
}

func (h Handlers) Create(ctx *gin.Context) {

	var cmd campaing.CommandCampaign

	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		foundation.AbortWithStatusError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Create(cmd)

	if err != nil {
		foundation.AbortWithStatusError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	foundation.Success(ctx, http.StatusCreated, cmd)
	return
}
