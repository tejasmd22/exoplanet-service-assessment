package stores

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource"
	"gofr.dev/pkg/gofr/http"

	"github.com/tejasmd22/exoplanet-service/filters"
	"github.com/tejasmd22/exoplanet-service/models"
)

type exoplanet struct{}

func New() *exoplanet {
	return &exoplanet{}
}

func (e *exoplanet) Create(ctx *gofr.Context, exoplanetCreateRequest *models.ExoplanetCreateRequest) (*models.Exoplanet, error) {
	query, values := e.generateInsertQuery(exoplanetCreateRequest)

	result, err := ctx.SQL.ExecContext(ctx, query, values...)
	if err != nil {
		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	exoplanetID, err := result.LastInsertId()
	if err != nil {
		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	return e.GetByID(ctx, int(exoplanetID))
}

func (e *exoplanet) GetByID(ctx *gofr.Context, id int) (*models.Exoplanet, error) {
	query := selectQuery + "WHERE id = ? "

	var exoplanet models.Exoplanet

	err := ctx.SQL.QueryRowContext(ctx, query, id).Scan(&exoplanet.ID, &exoplanet.Name,
		&exoplanet.Description, &exoplanet.Distance, &exoplanet.Radius, &exoplanet.Mass, &exoplanet.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.ErrorEntityNotFound{
				Name:  "exoplanet",
				Value: strconv.Itoa(id),
			}
		}

		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	return &exoplanet, nil
}

func (e *exoplanet) GetAll(ctx *gofr.Context, filter *filters.Exoplanet, page *models.Page) ([]*models.Exoplanet, error) {
	where, values := e.generateWhereClause(filter)

	query := selectQuery + where

	rows, err := ctx.SQL.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	defer rows.Close()

	var exoplanets []*models.Exoplanet

	for rows.Next() {
		var exoplanet models.Exoplanet

		err = rows.Scan(&exoplanet.ID, &exoplanet.Name, &exoplanet.Description, &exoplanet.Distance, &exoplanet.Radius, &exoplanet.Mass, &exoplanet.Type)
		if err != nil {
			return nil, datasource.ErrorDB{Err: err, Message: "error from db"}

		}

		exoplanets = append(exoplanets, &exoplanet)
	}

	return exoplanets, nil
}

func (e *exoplanet) Count(ctx *gofr.Context, filter *filters.Exoplanet) (int, error) {
	where, values := e.generateWhereClause(filter)

	query := countQuery + where

	var count int

	err := ctx.SQL.QueryRowContext(ctx, query, values...).Scan(&count)
	if err != nil {
		return 0, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	return count, nil
}

func (e *exoplanet) Update(ctx *gofr.Context, exoplanetUpdateRequest *models.ExoplanetUpdateRequest) (*models.Exoplanet, error) {
	setClause, values := e.generateUpdateQuery(exoplanetUpdateRequest)

	values = append(values, exoplanetUpdateRequest.ID)

	query := updateQuery + setClause + " WHERE id = ?"

	_, err := ctx.SQL.ExecContext(ctx, query, values...)
	if err != nil {
		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	return e.GetByID(ctx, exoplanetUpdateRequest.ID)
}

func (e *exoplanet) Delete(ctx *gofr.Context, id int) error {
	query := deleteQuery + "WHERE id = ?"

	_, err := ctx.SQL.ExecContext(ctx, query, id)
	if err != nil {
		return datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	return nil
}
