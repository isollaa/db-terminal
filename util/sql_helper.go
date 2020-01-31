package util

import (
	"database/sql"
	"fmt"

	"github.com/isollaa/dbterm"
)

var ListSQL = map[string]SQL{}

type SQL interface {
	DSNFormat() string
	GetQueryDB() string
	GetQueryTable() string
	GetDiskSpace(config dbterm.Config) (map[string]string, error)
}

func SQLDial(config dbterm.Config) (*sql.DB, error) {
	t := config[dbterm.DRIVER].(string)
	dsn := fmt.Sprintf(ListSQL[t].DSNFormat(),
		config[dbterm.USERNAME],
		config[dbterm.PASSWORD],
		config[dbterm.HOST],
		config[dbterm.PORT],
		config[dbterm.DBNAME],
	)

	db, err := sql.Open(t, dsn)
	if err != nil {
		return nil, err
	}

	// add conn setting

	return db, nil
}

func GetSQLListSession(config dbterm.Config) string {
	t := config[dbterm.DRIVER].(string)
	switch config[dbterm.FLAG_STAT] {
	case dbterm.FLAG_DB:
		return ListSQL[t].GetQueryDB()
	case dbterm.FLAG_COLL:
		return ListSQL[t].GetQueryTable()
	}
	return ""
}

func GetSQLDiskQuery(config dbterm.Config) (map[string]string, error) {
	t := config[dbterm.DRIVER].(string)
	return ListSQL[t].GetDiskSpace(config)
}
