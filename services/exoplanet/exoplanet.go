package services

import (
	"gofr.dev/pkg/gofr"

	"github.com/tejasmd22/exoplanet-service/filters"
	"github.com/tejasmd22/exoplanet-service/models"
	"github.com/tejasmd22/exoplanet-service/services"
	"github.com/tejasmd22/exoplanet-service/stores"
)

type exoplanet struct {
	exoplanetStores stores.Exoplanet
}

func New(exoplanetStores stores.Exoplanet) services.Exoplanet {
	return &exoplanet{exoplanetStores: exoplanetStores}
}

func (e *exoplanet) Create(ctx *gofr.Context, exoplanetCreateRequest *models.ExoplanetCreateRequest) (*models.Exoplanet, error) {
	exoplanet, err := e.exoplanetStores.Create(ctx, exoplanetCreateRequest)
	if err != nil {
		return nil, err
	}

	return exoplanet, nil
}

func (e *exoplanet) GetByID(ctx *gofr.Context, id int) (*models.Exoplanet, error) {
	exoplanet, err := e.exoplanetStores.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return exoplanet, nil
}

func (e *exoplanet) GetAll(ctx *gofr.Context, filter *filters.Exoplanet, page *models.Page) ([]*models.Exoplanet, error) {
	exoplanets, err := e.exoplanetStores.GetAll(ctx, filter, page)
	if err != nil {
		return nil, err
	}

	return exoplanets, nil
}

func (e *exoplanet) Count(ctx *gofr.Context, filter *filters.Exoplanet) (int, error) {
	count, err := e.exoplanetStores.Count(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (e *exoplanet) Update(ctx *gofr.Context, exoplanetUpdateRequest *models.ExoplanetUpdateRequest) (*models.Exoplanet, error) {
	_, err := e.GetByID(ctx, exoplanetUpdateRequest.ID)
	if err != nil {
		return nil, err
	}

	exoplanet, err := e.exoplanetStores.Update(ctx, exoplanetUpdateRequest)
	if err != nil {
		return nil, err
	}

	return exoplanet, nil
}

func (e *exoplanet) Delete(ctx *gofr.Context, id int) error {
	_, err := e.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = e.exoplanetStores.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *exoplanet) CalculateFuelCost(ctx *gofr.Context, id, crewCapacity int) (float64, error) {
	exoplanet, err := e.GetByID(ctx, id)
	if err != nil {
		return 0, err
	}

	var gravity float64

	if exoplanet.Type == models.GasGiant {
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else if exoplanet.Type == models.Terrestrial {
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}

	fuelCost := exoplanet.Distance / (gravity * gravity) * float64(crewCapacity)

	return fuelCost, nil
}
