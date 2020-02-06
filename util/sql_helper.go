package util

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/isollaa/dbterm/config"
	_ "github.com/lib/pq"
)

var ListSQL = map[string]SQL{}

type SQL interface {
	DSNFormat() string
	GetQueryDB() string
	GetQueryTable() string
	GetDiskSpace(c config.Config) (map[string]string, error)
}

func SQLDial(c config.Config) (*sql.DB, error) {
	// t := c[config.DRIVER].(string)
	t := c[config.DRIVER].(string)
	l, supported := ListSQL[t]
	if !supported {
		fmt.Printf("Error: Command not supported on selected database: %s \n", c[config.DRIVER])
		os.Exit(1)
	}
	dsn := fmt.Sprintf(l.DSNFormat(),
		c[config.USERNAME],
		c[config.PASSWORD],
		c[config.HOST],
		c[config.PORT],
		c[config.DBNAME],
	)

	db, err := sql.Open(t, dsn)
	if err != nil {
		return nil, err
	}

	// add conn setting

	return db, nil
}

func GetSQLListSession(c config.Config) string {
	t := c[config.DRIVER].(string)
	switch c[config.FLAG_STAT] {
	case config.FLAG_DB:
		return ListSQL[t].GetQueryDB()
	case config.FLAG_COLL:
		return ListSQL[t].GetQueryTable()
	}
	return ""
}

func GetSQLDiskQuery(c config.Config) (map[string]string, error) {
	t := c[config.DRIVER].(string)
	return ListSQL[t].GetDiskSpace(c)
}
