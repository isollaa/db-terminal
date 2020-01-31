package sql

import (
	"errors"
	"fmt"

	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type mysql struct{}

func (m *mysql) Autofill(config dbterm.Config) {
	if config[dbterm.PORT] == 0 {
		config[dbterm.PORT] = 3306
	}
	if config[dbterm.DBNAME] == "" {
		config[dbterm.DBNAME] = "mqtt"
	}
	if config[dbterm.USERNAME] == "" {
		config[dbterm.USERNAME] = "root"
	}
}

func (m *mysql) DSNFormat() string {
	return "%s:%s@tcp(%s:%d)/%s"
}

func (m *mysql) GetQueryDB() string {
	return "SHOW DATABASES"
}

func (m *mysql) GetQueryTable() string {
	return "SHOW TABLES"
}

func (m *mysql) GetDiskSpace(config dbterm.Config) (map[string]string, error) {
	v := map[string]string{}
	switch config[dbterm.FLAG_STAT] {
	case "db":
		return v, errors.New("disk status not available")
	case "coll":
		v["title"] = fmt.Sprintf("Table - %s", config[dbterm.COLLECTION])
		v["query"] = fmt.Sprintf("SELECT (data_length+index_length)/power(1024,1) FROM information_schema.tables WHERE table_schema='%s' and table_name='%s'", config[dbterm.DBNAME], config[dbterm.COLLECTION])
	default:
		return nil, fmt.Errorf("no such command: '%s'", config[dbterm.FLAG_STAT])
	}
	return v, nil
}

func init() {
	util.ListSQL["mysql"] = &mysql{}
}
