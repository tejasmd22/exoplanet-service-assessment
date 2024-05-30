package main

import (
	"gofr.dev/pkg/gofr"

	exoplanetHTTP "github.com/tejasmd22/exoplanet-service/http/exoplanet"
	"github.com/tejasmd22/exoplanet-service/migrations"
	exoplanetService "github.com/tejasmd22/exoplanet-service/services/exoplanet"
	exoplanetStore "github.com/tejasmd22/exoplanet-service/stores/exoplanet"
)

func main() {
	app := gofr.New()

	exoplanetStore := exoplanetStore.New()
	exoplanetService := exoplanetService.New(exoplanetStore)
	exoplanetHTTP := exoplanetHTTP.New(exoplanetService)

	app.Migrate(migrations.All())

	app.POST("/exoplanet", exoplanetHTTP.Create)
	app.GET("/exoplanet", exoplanetHTTP.GetAll)
	app.GET("/exoplanet/{id}", exoplanetHTTP.GetByID)
	app.PUT("/exoplanet/{id}", exoplanetHTTP.Update)
	app.DELETE("/exoplanet/{id}", exoplanetHTTP.Delete)
	app.GET("/exoplanet/{id}/fuel-estimation", exoplanetHTTP.CalculateFuelCost)

	app.Run()
}
