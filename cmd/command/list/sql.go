package list

import (
	_ "github.com/go-sql-driver/mysql"
	s "github.com/isollaa/db-terminal/cmd/init/sql"
	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
)

func sql(c t.Config, svc registry.Initial) error {
	ss := svc.(*s.SQL)
	rows, err := ss.Session.Query(s.GetListSession(c))
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
	ss.Result = result
	return nil
}

func init() {
	register("sql", sql)
}
