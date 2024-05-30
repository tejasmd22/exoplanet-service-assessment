package http

import (
	"strconv"
	"sync"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	pageHTTP "github.com/tejasmd22/exoplanet-service/http"
	"github.com/tejasmd22/exoplanet-service/models"
	"github.com/tejasmd22/exoplanet-service/services"
)

type handler struct {
	exoplanetServices services.Exoplanet
}

func New(exoplanetServices services.Exoplanet) *handler {
	return &handler{exoplanetServices: exoplanetServices}
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var exoplanetCreateRequest models.ExoplanetCreateRequest

	err := ctx.Bind(&exoplanetCreateRequest)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"request body"}}
	}

	if err := exoplanetCreateRequest.Validate(); err != nil {
		return nil, err
	}

	exoplanet, err := h.exoplanetServices.Create(ctx, &exoplanetCreateRequest)
	if err != nil {
		return nil, err
	}

	return models.ExoplanetResponse{
		Exoplanet: exoplanet,
	}, nil
}

func (h *handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	filter, err := h.validateAndGetFilters(ctx)
	if err != nil {
		return nil, err
	}

	pageNumber := ctx.Param("page")
	perPage := ctx.Param("perPage")
	paginated := ctx.Param("paginated")

	page, err := pageHTTP.ParsePagination(pageNumber, perPage, paginated)
	if err != nil {
		return nil, err
	}

	var (
		wg         sync.WaitGroup
		errChan    = make(chan error, 2)
		exoplanets []*models.Exoplanet
		count      int
	)

	wg.Add(2)

	go func() {
		defer wg.Done()

		exoplanets, err = h.exoplanetServices.GetAll(ctx, filter, page)
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()

		count, err = h.exoplanetServices.Count(ctx, filter)
		if err != nil {
			errChan <- err
		}
	}()

	wg.Wait()

	close(errChan)

	for err = range errChan {
		if err != nil {
			return nil, err
		}
	}

	if !page.Paginated {
		page.Number = pageHTTP.PerPageLowerLimit
		page.PerPage = count
	}

	return models.ExoplanetDetail{
		Data:   exoplanets,
		Count:  count,
		Limit:  page.PerPage,
		Offset: (page.Number - 1) * page.PerPage,
	}, nil
}

func (h *handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	exoplanetID, err := h.validateExoplanetID(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	exoplanet, err := h.exoplanetServices.GetByID(ctx, exoplanetID)
	if err != nil {
		return nil, err
	}

	return models.ExoplanetResponse{Exoplanet: exoplanet}, nil
}

func (h *handler) Update(ctx *gofr.Context) (interface{}, error) {
	exoplanetID, err := h.validateExoplanetID(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	var exoplanetUpdateRequest models.ExoplanetUpdateRequest

	err = ctx.Bind(&exoplanetUpdateRequest)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"request body"}}
	}

	if err := exoplanetUpdateRequest.Validate(); err != nil {
		return nil, err
	}

	exoplanetUpdateRequest.ID = exoplanetID

	exoplanet, err := h.exoplanetServices.Update(ctx, &exoplanetUpdateRequest)
	if err != nil {
		return nil, err
	}

	return models.ExoplanetResponse{Exoplanet: exoplanet}, nil
}

func (h *handler) Delete(ctx *gofr.Context) (interface{}, error) {
	exoplanetID, err := h.validateExoplanetID(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	err = h.exoplanetServices.Delete(ctx, exoplanetID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handler) CalculateFuelCost(ctx *gofr.Context) (interface{}, error) {
	exoplanetID, err := h.validateExoplanetID(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	crewCapacityStr := ctx.Param("crewCapacity")
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil || crewCapacity < 1 {
		return nil, http.ErrorInvalidParam{Params: []string{"crewCapacity"}}
	}

	fuelCost, err := h.exoplanetServices.CalculateFuelCost(ctx, exoplanetID, crewCapacity)
	if err != nil {
		return nil, err
	}

	return struct {
		FuelCost float64 `json:"fuelCost"`
	}{FuelCost: fuelCost}, nil
}
