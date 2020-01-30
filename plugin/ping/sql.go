package ping

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type sql struct{}

func (m *sql) Ping(config dbterm.Config) error {
	db, err := util.SQLDial(config)
	if err != nil {
		return err
	}

	return db.Ping()
}

func init() {
	supportedDB["sql"] = &sql{}
}
