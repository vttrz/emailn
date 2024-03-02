package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vttrz/emailn/internal"
)

func MapRoutes(router *gin.Engine, application *internal.Application) {

	v1 := router.Group("/v1")
	{
		campaign := v1.Group("/campaign")
		{
			campaign.GET("", application.CampaignController.List)
			campaign.POST("", application.CampaignController.Create)

		}
	}
}
