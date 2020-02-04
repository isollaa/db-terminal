package disk

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
	query, err := util.GetMongoDiskQuery(c)
	if err != nil {
		return err
	}
	err = session.DB(c[config.DBNAME].(string)).Run(query, &r.Value)
	if err != nil {
		return err
	}
	return nil
}

// func init() {
// 	m := &Mongo{}
// 	plugin.RegisterDriver("mongo", m.disk)
// }
