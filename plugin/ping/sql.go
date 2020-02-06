package ping

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/util"
	_ "github.com/lib/pq"
)

func SQL(r *registry.Result, c config.Config) error {
	start := time.Now()
	defer func() {
		r.Value = fmt.Sprintf("Ping done in %d ms", time.Now().Sub(start).Microseconds())
	}()

	session, err := util.SQLDial(c)
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Ping()
}

// func init() {
// 	m := &SQL{}
// 	registry.RegisterDriver("sql", m.ping)
// }
