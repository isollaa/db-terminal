package list

import (
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type sql struct {
	Result interface{}
}

func (s *sql) List(config dbterm.Config) error {
	session, err := util.SQLDial(config)
	if err != nil {
		return err
	}
	rows, err := session.Query(util.GetSQLListSession(config))
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
	s.Result = result
	return nil
}

func init() {
	supportedDB["sql"] = &sql{}
}
