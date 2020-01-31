package disk

import (
	"errors"
	"fmt"

	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type sql struct {
	Result interface{}
}

func (s *sql) Disk(config dbterm.Config) error {
	session, err := util.SQLDial(config)
	if err != nil {
		return err
	}
	v, err := util.GetSQLDiskQuery(config)
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
	if config[dbterm.DRIVER] == "mysql" {
		table = table + " kB"
	}
	s.Result = fmt.Sprintf("%s, Disk Size: %s", v["title"], table)
	return nil
}

func init() {
	supportedDB["sql"] = &sql{}
}
