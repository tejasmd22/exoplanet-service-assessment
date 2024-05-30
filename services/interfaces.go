package services

import (
	"gofr.dev/pkg/gofr"

	"github.com/tejasmd22/exoplanet-service/filters"
	"github.com/tejasmd22/exoplanet-service/models"
)

type Exoplanet interface {
	Create(ctx *gofr.Context, exoplanetCreateRequest *models.ExoplanetCreateRequest) (*models.Exoplanet, error)
	GetByID(ctx *gofr.Context, id int) (*models.Exoplanet, error)
	GetAll(ctx *gofr.Context, filter *filters.Exoplanet, page *models.Page) ([]*models.Exoplanet, error)
	Count(ctx *gofr.Context, filter *filters.Exoplanet) (int, error)
	Update(ctx *gofr.Context, exoplanetUpdateRequest *models.ExoplanetUpdateRequest) (*models.Exoplanet, error)
	Delete(ctx *gofr.Context, id int) error
	CalculateFuelCost(ctx *gofr.Context, id, crewCapacity int) (float64, error)
}
