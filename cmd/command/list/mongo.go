package list

import (
	m "github.com/isollaa/db-terminal/cmd/init/mongo"
	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
)

func mongo(c t.Config, svc registry.Initial) error {
	s := svc.(*m.Mongo)
	result, err := m.GetListSession(c, s)
	if err != nil {
		return err
	}
	s.Result = result
	return nil
}

func init() {
	register("mongo", mongo)
}
