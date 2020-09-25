package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

// migrations contains the queries needed to construct the database schema.
// Entries should never be removed from this slice once they have been ran in
// production.
//
// Including the queries directly in this file has the same pros/cons mentioned
// in seeds.go

var (
	migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Add users",
			Script: `
					CREATE TABLE IF NOT EXISTS users(
						id BIGINT(20) NOT NULL AUTO_INCREMENT,
						first_name VARCHAR(45) NULL,
						last_name VARCHAR(45) NULL,
						email VARCHAR(45) NOT NULL,
						status VARCHAR(45) NOT NULL,
						password VARCHAR(32) NOT NULL,
						date_created DATETIME NOT NULL,
						date_updated DATETIME NOT NULL,
						PRIMARY KEY(id),
						UNIQUE INDEX email_UNIQUE (email ASC)
					);
			`,
		},
	}
)

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sql.DB) error {

	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
