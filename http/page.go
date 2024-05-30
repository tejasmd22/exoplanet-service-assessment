package http

import (
	"strconv"

	"gofr.dev/pkg/gofr/http"

	"github.com/tejasmd22/exoplanet-service/models"
)

const (
	pageLimit         = 20
	PerPageLowerLimit = 1
	perPageUpperLimit = 100
)

func ParsePagination(pageNumber, perPage, paginated string) (*models.Page, error) {
	var page = &models.Page{}

	paginatedBool, err := validatePaginated(paginated)
	if err != nil {
		return nil, err
	}

	if !paginatedBool {
		page.Paginated = paginatedBool

		return page, nil
	}

	pageNumInt, err := validatePageNumber(pageNumber)
	if err != nil {
		return nil, err
	}

	perPageInt, err := validatePerPage(perPage)
	if err != nil {
		return nil, err
	}

	page.Paginated = paginatedBool
	page.Number = pageNumInt
	page.PerPage = perPageInt

	return page, nil
}

func validatePaginated(paginated string) (bool, error) {
	if paginated == "" {
		return true, nil
	}

	paginatedBool, err := strconv.ParseBool(paginated)
	if err != nil {
		return false, http.ErrorInvalidParam{Params: []string{"paginate"}}
	}

	return paginatedBool, nil
}

func validatePageNumber(pageNumber string) (int, error) {
	if pageNumber == "" {
		pageNumber = "1"
	}

	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 1 {
		return 0, http.ErrorInvalidParam{Params: []string{"page"}}
	}

	return pageNumberInt, err
}

func validatePerPage(perPage string) (int, error) {
	if perPage == "" {
		return pageLimit, nil
	}

	perPageInt, err := strconv.Atoi(perPage)
	if err != nil || perPageInt < PerPageLowerLimit || perPageInt > perPageUpperLimit {
		return 0, http.ErrorInvalidParam{Params: []string{"perPage"}}
	}

	return perPageInt, nil
}
