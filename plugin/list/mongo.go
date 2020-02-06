package list

import (
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/util"
)

func Mongo(r *registry.Result, c config.Config) error {
	session, err := util.MongoDial(c)
	if err != nil {
		return err
	}
	defer session.Close()
	r.Value, err = util.GetMongoListSession(c, session)
	if err != nil {
		return err
	}

	return nil
}

// func init() {
// 	m := &Mongo{}
// 	plugin.RegisterDriver("mongo", m.list)
// }
