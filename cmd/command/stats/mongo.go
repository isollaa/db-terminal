package stats

import (
	"github.com/globalsign/mgo/bson"
	m "github.com/isollaa/db-terminal/cmd/init/mongo"
	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
)

func mongo(c t.Config, info string, svc registry.Initial) error {
	var query interface{}
	s := svc.(*m.Mongo)
	switch info {
	case "db":
		query = "dbstats"
	case "coll":
		query = &bson.D{bson.DocElem{"collstats", s.Collection}}
	}
	result := bson.M{}
	err := s.Session.DB(s.DBName).Run(query, &result)
	if err != nil {
		return err
	}
	s.Result = result
	return nil
}

func init() {
	register("mongo", mongo)
}
