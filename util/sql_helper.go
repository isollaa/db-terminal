package util

import (
	"database/sql"
	"fmt"

	"github.com/isollaa/dbterm"
)

var ListSQL = map[string]SQL{}

type SQL interface {
	Autofill(dbterm.Config)
}

func SQLDial(config dbterm.Config) (*sql.DB, error) {
	t := config[dbterm.DB_DRIVER].(string)
	ListSQL[t].Autofill(config)
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=require",
		config[dbterm.DB_DRIVER],
		config[dbterm.DB_USERNAME],
		config[dbterm.DB_PASSWORD],
		config[dbterm.DB_HOST],
		config[dbterm.DB_PORT],
		config[dbterm.DB_DBNAME],
	)

	db, err := sql.Open(t, dsn)
	if err != nil {
		return nil, err
	}

	// add conn setting

	return db, nil
}
