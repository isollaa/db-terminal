package disk

import (
	"github.com/globalsign/mgo/bson"
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type mongo struct {
	Result interface{}
}

func (s *mongo) Disk(config dbterm.Config) error {
	session, err := util.MongoDial(config)
	if err != nil {
		return err
	}
	query, err := util.GetMongoDiskQuery(config)
	if err != nil {
		return err
	}
	result := bson.M{}
	err = session.DB(config[dbterm.DBNAME].(string)).Run(query, &result)
	if err != nil {
		return err
	}
	s.Result = result
	return nil
}
