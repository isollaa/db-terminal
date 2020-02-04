package ping

import (
	"fmt"
	"time"

	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/util"
)

func Mongo(r *registry.Result, c config.Config) error {
	start := time.Now()
	defer func() {
		r.Value = fmt.Sprintf("Ping done in %d ms", time.Now().Sub(start).Microseconds())
	}()

	db, err := util.MongoDial(c)
	if err != nil {
		return err
	}

	return db.Ping()
}

// func init() {
// 	m := &Mongo{}
// 	registry.RegisterDriver("mongo", m.ping)
// }
