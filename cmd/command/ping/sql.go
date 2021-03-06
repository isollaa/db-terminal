package ping

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	s "github.com/isollaa/dbterm/cmd/init/sql"
	"github.com/isollaa/dbterm/registry"
)

func sql(svc registry.Initial) error {
	s := svc.(*s.SQL)
	err := s.Session.Ping()
	if err != nil {
		return err
	}
	s.Result = fmt.Sprintf("-- %s server is ok.", s.Driver)
	return nil
}

func init() {
	register("sql", sql)
}
