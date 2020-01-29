package ping

import (
	"fmt"

	m "github.com/isollaa/db-terminal/cmd/db/mongo"
	"github.com/isollaa/db-terminal/registry"
)

func mongo(svc registry.Initial) error {
	s := svc.(*m.Mongo)
	err := s.Session.Ping()
	if err != nil {
		return err
	}
	s.Result = fmt.Sprintf("-- MongoDB server is ok.")
	return nil
}

func init() {
	register("mongo", mongo)
}
