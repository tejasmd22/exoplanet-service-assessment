package http

import (
	"strconv"
	"strings"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/tejasmd22/exoplanet-service/filters"
)

func (h *handler) validateAndGetFilters(ctx *gofr.Context) (*filters.Exoplanet, error) {
	var filter filters.Exoplanet

	name := ctx.Param("name")
	if name != "" {
		filter.Name = name
	}

	if ctx.Param("ids") != "" {
		ids := strings.Split(ctx.Param("ids"), ",")
		for _, id := range ids {
			exoplanetID, err := strconv.Atoi(id)
			if err != nil {
				return nil, http.ErrorInvalidParam{Params: []string{"ids"}}
			}

			filter.IDs = append(filter.IDs, exoplanetID)
		}
	}

	if ctx.Param("distance") != "" {
		distance, err := strconv.ParseFloat(ctx.Param("distance"), 64)
		if err != nil || distance < 1 {
			return nil, http.ErrorInvalidParam{Params: []string{"distance"}}
		}

		filter.Distance = distance
	}

	if ctx.Param("radius") != "" {
		radius, err := strconv.ParseFloat(ctx.Param("radius"), 64)
		if err != nil || radius < 1 {
			return nil, http.ErrorInvalidParam{Params: []string{"radius"}}
		}

		filter.Radius = radius
	}

	if ctx.Param("mass") != "" {
		mass, err := strconv.ParseFloat(ctx.Param("mass"), 64)
		if err != nil || mass < 1 {
			return nil, http.ErrorInvalidParam{Params: []string{"mass"}}
		}

		filter.Mass = mass
	}

	exoplanetType := ctx.Param("type")
	if exoplanetType != "" {
		if !strings.EqualFold(exoplanetType, "GasGiant") && !strings.EqualFold(exoplanetType, "Terrestrial") {
			return nil, http.ErrorInvalidParam{Params: []string{"type"}}
		}

		filter.Type = strings.ToUpper(exoplanetType)
	}

	return &filter, nil
}

func (h *handler) validateExoplanetID(id string) (int, error) {
	if id == "" {
		return 0, http.ErrorMissingParam{Params: []string{"id"}}
	}

	exoplanetID, err := strconv.Atoi(id)
	if err != nil {
		return 0, http.ErrorInvalidParam{Params: []string{"id"}}
	}

	return exoplanetID, nil
}
