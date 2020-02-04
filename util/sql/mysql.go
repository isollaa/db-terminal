// +build mysql
// build

package sql

import (
	"errors"
	"fmt"

	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type mysql struct{}

func (m *mysql) DSNFormat() string {
	return "%s:%s@tcp(%s:%d)/%s"
}

func (m *mysql) GetQueryDB() string {
	return "SHOW DATABASES"
}

func (m *mysql) GetQueryTable() string {
	return "SHOW TABLES"
}

func (m *mysql) GetDiskSpace(config config.Config) (map[string]string, error) {
	v := map[string]string{}
	switch config[config.FLAG_STAT] {
	case "db":
		return v, errors.New("disk status not available")
	case "coll":
		v["title"] = fmt.Sprintf("Table - %s", config[config.COLLECTION])
		v["query"] = fmt.Sprintf("SELECT (data_length+index_length)/power(1024,1) FROM information_schema.tables WHERE table_schema='%s' and table_name='%s'", config[config.DBNAME], config[config.COLLECTION])
	default:
		return nil, fmt.Errorf("no such command: '%s'", config[config.FLAG_STAT])
	}
	return v, nil
}

func init() {
	util.ListSQL["mysql"] = &mysql{}
}
