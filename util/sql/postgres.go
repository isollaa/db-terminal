package sql

import (
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type postgres struct{}

func (m *postgres) Autofill(config dbterm.Config) {
	if config[dbterm.DB_PORT] == 0 {
		config[dbterm.DB_PORT] = 5432
	}
	if config[dbterm.DB_DBNAME] == "" {
		config[dbterm.DB_DBNAME] = "postgres"
	}
	if config[dbterm.DB_USERNAME] == "" {
		config[dbterm.DB_USERNAME] = "postgres"
	}
	if config[dbterm.DB_PASSWORD] == "" {
		config[dbterm.DB_PASSWORD] = "12345"
	}
}

func init() {
	util.ListSQL["postgres"] = &mysql{}
}
