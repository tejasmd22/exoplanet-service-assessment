package migrations

import "gofr.dev/pkg/gofr/migration"

const createExoplanet = "CREATE TABLE IF NOT EXISTS exoplanet (" +
	"id int(11) NOT NULL AUTO_INCREMENT, " +
	"name varchar(200), " +
	"description TEXT, " +
	"distance float NOT NULL, " +
	"radius float NOT NULL, " +
	"mass float NOT NULL, " +
	"type varchar(100), " +
	"created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
	"updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, " +
	"deleted_at timestamp NULL DEFAULT NULL, " +
	"PRIMARY KEY (id) " +
	");"

func createExoplanetTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createExoplanet)
			if err != nil {
				return err
			}
			
			return nil
		},
	}
}
