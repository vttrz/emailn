package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	campaing "github.com/vttrz/emailn/internal/domain/campaign"
	"github.com/vttrz/emailn/internal/infrastrucutre/controller"
)

func Routers() {
	r := gin.Default()

	prefix := "/campaigns"

	repo := campaing.NewRepository(nil)

	s := campaing.NewService(repo)
	h := controller.NewHandlers(s, repo)

	r.POST(prefix, h.Create)

	err := r.Run(":8080")

	if err != nil {
		return
	}

	fmt.Println("application running")
}
