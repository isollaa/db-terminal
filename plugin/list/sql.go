package list

import (
	"fmt"

	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/plugin"
	"github.com/isollaa/dbterm/util"
)

func SQL(r *registry.Result, c config.Config) error {
	session, err := util.SQLDial(c)
	if err != nil {
		return err
	}
	query := util.GetSQLListSession(c)
	if query == "" {
		return fmt.Errorf("Error: unable to get query")
	}
	rows, err := session.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	result := []string{}
	for rows.Next() {
		res := ""
		rows.Scan(&res)
		result = append(result, res)
	}
	r.Value = result
	return nil
}

// func init() {
// 	m := &SQL{}
// 	plugin.RegisterDriver("sql", m.list)
// }
