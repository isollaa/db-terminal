package sql

import (
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type mysql struct{}

func (m *mysql) Autofill(config dbterm.Config) {
	if config[dbterm.DB_PORT] == 0 {
		config[dbterm.DB_PORT] = 3306
	}
	if config[dbterm.DB_DBNAME] == "" {
		config[dbterm.DB_DBNAME] = "mqtt"
	}
	if config[dbterm.DB_USERNAME] == "" {
		config[dbterm.DB_USERNAME] = "root"
	}
}

func init() {
	util.ListSQL["mysql"] = &mysql{}
}
