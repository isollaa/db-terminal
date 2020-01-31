package sql

import (
	"fmt"

	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type postgres struct{}

func (m *postgres) Autofill(config dbterm.Config) {
	if config[dbterm.PORT] == 0 {
		config[dbterm.PORT] = 5432
	}
	if config[dbterm.DBNAME] == "" {
		config[dbterm.DBNAME] = "postgres"
	}
	if config[dbterm.USERNAME] == "" {
		config[dbterm.USERNAME] = "postgres"
	}
	if config[dbterm.PASSWORD] == "" {
		config[dbterm.PASSWORD] = "12345"
	}
}

func (m *postgres) DSNFormat() string {
	return "postgres://%s:%s@%s:%d/%s?sslmode=require"
}

func (m *postgres) GetQueryDB() string {
	return "SELECT datname FROM pg_database WHERE datistemplate = false"
}

func (m *postgres) GetQueryTable() string {
	return "SELECT table_schema,table_name FROM information_schema.tables ORDER BY table_schema,table_name"
}

func (m *postgres) GetDiskSpace(config dbterm.Config) (map[string]string, error) {
	v := map[string]string{}
	info := ""
	attrib := ""
	switch config[dbterm.FLAG_STAT] {
	case "db":
		info = "pg_database_size"
		attrib = config[dbterm.DBNAME].(string)
		v["title"] = "DB - " + attrib
	case "coll":
		info = "pg_total_relation_size"
		attrib = config[dbterm.COLLECTION].(string)
		v["title"] = fmt.Sprintf("Table - %s", attrib)
	default:
		return nil, fmt.Errorf("no such command: '%s'", config[dbterm.FLAG_STAT])
	}
	v["query"] = fmt.Sprintf("SELECT pg_size_pretty(%s('%s'))", info, attrib)
	return v, nil
}

func init() {
	util.ListSQL["postgres"] = &postgres{}
}
