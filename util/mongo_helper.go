package util

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/isollaa/dbterm"
)

func MongoDial(config dbterm.Config) (*mgo.Session, error) {
	if config[dbterm.DB_PORT] == 0 {
		config[dbterm.DB_PORT] = 27017
	}
	if config[dbterm.DB_DBNAME] == "" {
		config[dbterm.DB_DBNAME] = "xsaas_ctms"
	}
	dsn := fmt.Sprintf("mongodb://%s:%d/%s",
		config[dbterm.DB_HOST],
		config[dbterm.DB_PORT],
		config[dbterm.DB_DBNAME],
	)

	session, err := mgo.Dial(dsn)
	if err != nil {
		return nil, err
	}

	// TODO: add session setting

	return session, nil
}
