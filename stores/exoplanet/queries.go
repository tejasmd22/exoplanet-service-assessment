package stores

import (
	"strings"

	"github.com/tejasmd22/exoplanet-service/filters"
	"github.com/tejasmd22/exoplanet-service/models"
)

const (
	insertQuery = "INSERT INTO exoplanet "
	selectQuery = "SELECT id, name, description, distance, radius, mass, type FROM exoplanet "
	countQuery  = "SELECT COUNT(*) FROM exoplanet "
	updateQuery = "UPDATE exoplanet "
	deleteQuery = "DELETE FROM exoplanet "
)

func (e *exoplanet) generateInsertQuery(exoplanetCreateRequest *models.ExoplanetCreateRequest) (query string, values []interface{}) {
	if exoplanetCreateRequest.Name != "" {
		query += "name, "

		values = append(values, exoplanetCreateRequest.Name)
	}

	if exoplanetCreateRequest.Description != "" {
		query += "description, "

		values = append(values, exoplanetCreateRequest.Description)
	}

	if exoplanetCreateRequest.Distance > 0 {
		query += "distance, "

		values = append(values, exoplanetCreateRequest.Distance)
	}

	if exoplanetCreateRequest.Radius > 0 {
		query += "radius, "

		values = append(values, exoplanetCreateRequest.Radius)
	}

	if exoplanetCreateRequest.Mass > 0 {
		query += "mass, "

		values = append(values, exoplanetCreateRequest.Mass)
	}

	if exoplanetCreateRequest.Type != "" {
		query += "type, "

		values = append(values, exoplanetCreateRequest.Type)
	}

	query = strings.Trim(query, ", ")
	if len(values) > 0 {
		query = "(" + query + ") VALUES"
		q := strings.Repeat("?, ", len(values))
		q = strings.Trim(q, ", ")
		query += "(" + q + ")"
	}

	query = insertQuery + query

	return query, values
}

func (e *exoplanet) generateWhereClause(filter *filters.Exoplanet) (where string, values []interface{}) {
	where += "WHERE deleted_at IS NULL AND "

	if len(filter.IDs) > 0 {
		q := strings.Trim(strings.Repeat("?, ", len(filter.IDs)), ", ")

		where += "id IN (" + q + ") AND "

		for _, userID := range filter.IDs {
			values = append(values, userID)
		}
	}

	if filter.Name != "" {
		where += "name = ? AND "

		values = append(values, filter.Name)
	}

	if filter.Distance > 0 {
		where += "distance = ? AND "

		values = append(values, filter.Distance)
	}

	if filter.Radius > 0 {
		where += "radius = ? AND "

		values = append(values, filter.Radius)
	}

	if filter.Mass > 0 {
		where += "mass = ? AND "

		values = append(values, filter.Mass)
	}

	if filter.Type != "" {
		where += "type = ? AND "

		values = append(values, filter.Type)
	}

	where = strings.TrimSuffix(where, "AND ")

	return
}

func (e *exoplanet) generateUpdateQuery(exoplanetUpdateRequest *models.ExoplanetUpdateRequest) (setClause string, values []interface{}) {
	setClause += "SET "

	if exoplanetUpdateRequest.Name != "" {
		setClause += "name = ?, "

		values = append(values, exoplanetUpdateRequest.Name)
	}

	if exoplanetUpdateRequest.Description != "" {
		setClause += "description = ?, "

		values = append(values, exoplanetUpdateRequest.Description)
	}

	if exoplanetUpdateRequest.Distance > 0 {
		setClause += "distance = ?, "

		values = append(values, exoplanetUpdateRequest.Distance)
	}

	if exoplanetUpdateRequest.Radius > 0 {
		setClause += "radius = ?, "

		values = append(values, exoplanetUpdateRequest.Radius)
	}

	if exoplanetUpdateRequest.Mass > 0 {
		setClause += "mass = ?, "

		values = append(values, exoplanetUpdateRequest.Mass)
	}

	if exoplanetUpdateRequest.Type != "" {
		setClause += "type = ?, "

		values = append(values, exoplanetUpdateRequest.Type)
	}

	setClause = strings.TrimSuffix(setClause, ", ")

	return
}
