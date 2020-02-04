package info

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/util"
)

func Mongo(r *registry.Result, c config.Config) error {
	session, err := util.MongoDial(c)
	if err != nil {
		return err
	}
	query := &bson.D{}
	info := fmt.Sprintf("%sInfo", c[config.FLAG_STAT])
	if info == "serverInfo" {
		query = &bson.D{{"serverStatus", 1}, {"repl", 0}, {"metrics", 0}, {"locks", 0}}
	} else {
		query = &bson.D{{info, 1}}
	}
	err = session.DB("admin").Run(query, &r.Value)
	if err != nil {
		return err
	}

	return nil
}

// func init() {
// 	m := &Mongo{}
// 	registry.RegisterDriver("mongo", m.info)
// }
