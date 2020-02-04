// +build postgres
// build

package sql

import (
	"fmt"

	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type postgres struct{}

func (m *postgres) DSNFormat() string {
	return "postgres://%s:%s@%s:%d/%s?sslmode=require"
}

func (m *postgres) GetQueryDB() string {
	return "SELECT datname FROM pg_database WHERE datistemplate = false"
}

func (m *postgres) GetQueryTable() string {
	return "SELECT table_schema,table_name FROM information_schema.tables ORDER BY table_schema,table_name"
}

func (m *postgres) GetDiskSpace(config config.Config) (map[string]string, error) {
	v := map[string]string{}
	info := ""
	attrib := ""
	switch config[config.FLAG_STAT] {
	case "db":
		info = "pg_database_size"
		attrib = config[config.DBNAME].(string)
		v["title"] = "DB - " + attrib
	case "coll":
		info = "pg_total_relation_size"
		attrib = config[config.COLLECTION].(string)
		v["title"] = fmt.Sprintf("Table - %s", attrib)
	default:
		return nil, fmt.Errorf("no such command: '%s'", config[config.FLAG_STAT])
	}
	v["query"] = fmt.Sprintf("SELECT pg_size_pretty(%s('%s'))", info, attrib)
	return v, nil
}

func init() {
	util.ListSQL["postgres"] = &postgres{}
}
