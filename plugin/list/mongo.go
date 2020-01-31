package list

import (
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type mongo struct {
	Result interface{}
}

func (s *mongo) List(config dbterm.Config) error {
	session, err := util.MongoDial(config)
	if err != nil {
		return err
	}
	result, err := util.GetMongoListSession(config, session)
	if err != nil {
		return err
	}
	s.Result = result

	return nil
}
