package ping

import (
	"testing"

	"github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	_ "github.com/isollaa/dbterm/util/sql/mysql"
	_ "github.com/isollaa/dbterm/util/sql/postgres"

	_ "github.com/isollaa/dbterm/util/sql/mysql"

	_ "github.com/isollaa/dbterm/util/sql/postgres"
)

type testCase struct {
	config.Config
	Valid bool
}

var testCases = []testCase{
	{
		Config: config.Config{
			"driver":     "mongo",
			"host":       "localhost",
			"port":       27017,
			"username":   "",
			"password":   "",
			"dbname":     "xsaas_ctms",
			"collection": "relationship",
			"category":   "mongo",
			"stat":       "",
			"beauty":     true,
			"prompt":     false,
		},
		Valid: true,
	},
	{
		Config: config.Config{
			"driver":     "mysql",
			"host":       "localhost",
			"port":       3306,
			"username":   "root",
			"password":   "",
			"dbname":     "mqtt",
			"collection": "ListCLient",
			"category":   "sql",
			"stat":       "",
			"beauty":     true,
			"prompt":     false,
		},
		Valid: true,
	},
	{
		Config: config.Config{
			"driver":     "postgres",
			"host":       "localhost",
			"port":       5432,
			"username":   "postgres",
			"password":   "12345",
			"dbname":     "postgres",
			"collection": "listclient",
			"category":   "sql",
			"stat":       "",
			"beauty":     true,
			"prompt":     false,
		},
		Valid: true,
	},
}

func TestCommand(t *testing.T) {
	registry.RegisterDriver(Mongo)
	registry.RegisterDriver(SQL)
	for i, v := range testCases {
		cat := v.Config["category"].(string)
		if command, supported := registry.Driver(cat, helper.GetName(helper.PACKAGE, command)); supported {
			r := registry.Result{}
			err := command(&r, v.Config)
			valid := err == nil
			if valid != v.Valid {
				t.Error("Error: ", i, err)
			}
		}
	}
}
