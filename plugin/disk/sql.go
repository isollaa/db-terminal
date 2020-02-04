package disk

import (
	"errors"
	"fmt"

	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/util"
)

func Sql(r *registry.Result, c config.Config) error {
	session, err := util.SQLDial(c)
	if err != nil {
		return err
	}
	v, err := util.GetSQLDiskQuery(c)
	if err != nil {
		return err
	}
	row, err := session.Query(v["query"])
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
	if c[config.DRIVER] == "mysql" {
		table = table + " kB"
	}
	r.Value = fmt.Sprintf("%s, Disk Size: %s", v["title"], table)
	return nil
}

// func init() {
// 	m := &SQL{}
// 	registry.RegisterDriver("sql", m.disk)
// }
