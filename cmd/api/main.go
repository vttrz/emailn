package main

import (
	"github.com/vttrz/emailn/internal"
	"github.com/vttrz/emailn/internal/infrastrucutre/routes"
)

func main() {

	// TODO: start db

	application := internal.NewApplication()
	defaultRouter := routes.DefaultRouter()

	routes.MapRoutes(defaultRouter, application)

	if err := defaultRouter.Run(":8080"); err != nil {
		panic(err)
	}
}
