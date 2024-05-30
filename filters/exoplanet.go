package filters

import (
	"fmt"
	"net/http"
)

type Exoplanet struct {
	IDs      []int
	Name     string
	Distance float64
	Radius   float64
	Mass     float64
	Type     string
}

type customError struct {
	error string
}

func (c customError) Error() string {
	return fmt.Sprintf("error: %s", c.error)
}

func (c customError) StatusCode() int {
	return http.StatusBadRequest
}
