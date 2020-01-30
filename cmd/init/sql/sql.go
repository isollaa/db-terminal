package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	_ "github.com/lib/pq"
)

type SQL struct {
	Driver     string
	DBName     string
	Collection string
	Result     interface{}
	Session    *sql.DB
}

func (m *SQL) AutoFill(c t.Config) {
	switch c[t.DB_DRIVER] {
	case "mysql":
		if c[t.DB_PORT] == 0 {
			c[t.DB_PORT] = 3306
		}
		if c[t.DB_DBNAME] == "" {
			c[t.DB_DBNAME] = "mqtt"
		}
		if c[t.DB_USERNAME] == "" {
			c[t.DB_USERNAME] = "root"
		}
	case "postgres":
		if c[t.DB_PORT] == 0 {
			c[t.DB_PORT] = 5432
		}
		if c[t.DB_DBNAME] == "" {
			c[t.DB_DBNAME] = "postgres"
		}
		if c[t.DB_USERNAME] == "" {
			c[t.DB_USERNAME] = "postgres"
		}
		if c[t.DB_PASSWORD] == "" {
			c[t.DB_PASSWORD] = "12345"
		}
	}
}

func (m *SQL) Connect(c t.Config) error {
	var err error
	m.Session, err = sql.Open(c[t.DB_DRIVER].(string), GetSource(c))
	if err != nil {
		return err
	}
	m.Driver = c[t.DB_DRIVER].(string)
	m.DBName = c[t.DB_DBNAME].(string)
	m.Collection = c[t.DB_COLLECTION].(string)
	return nil
}

func (m *SQL) Close() {
	defer m.Session.Close()
}

func init() {
	registry.RegisterDB(this)
}

func this() registry.Initial {
	return &SQL{}
}
