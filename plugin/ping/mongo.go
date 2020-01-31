package ping

import (
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type mongo struct{}

func (m *mongo) Ping(config dbterm.Config) error {
	db, err := util.MongoDial(config)
	if err != nil {
		return err
	}

	return db.Ping()
}
