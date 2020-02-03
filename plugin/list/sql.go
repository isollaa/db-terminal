package list

import (
	"fmt"

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
	query := util.GetSQLListSession(config)
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
	s.Result = result
	return nil
}
