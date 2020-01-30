package list

import (
	m "github.com/isollaa/dbterm/cmd/init/mongo"
	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
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
