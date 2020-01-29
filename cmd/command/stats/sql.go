package stats

import (
	"errors"
	"fmt"

	s "github.com/isollaa/db-terminal/cmd/init/sql"
	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
)

func sql(c t.Config, info string, svc registry.Initial) error {
	v, err := s.GetDiskSpace(info, c)
	if err != nil {
		return err
	}
	ss := svc.(*s.SQL)
	row, err := ss.Session.Query(v["query"])
	if err != nil {
		return err
	}
	defer row.Close()
	table := ""
	for row.Next() {
		row.Scan(&table)
	}
	if table == "" {
		return errors.New("data not found")
	}
	if ss.Driver == "mysql" {
		table = table + " kB"
	}
	ss.Result = fmt.Sprintf("%s, Disk Size: %s", v["title"], table)
	return nil
}

func init() {
	register("sql", sql)
}
