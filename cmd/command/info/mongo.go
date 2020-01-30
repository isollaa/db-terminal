package info

import (
	"github.com/globalsign/mgo/bson"
	m "github.com/isollaa/db-terminal/cmd/init/mongo"
	"github.com/isollaa/db-terminal/registry"
)

func mongo(info string, svc registry.Initial) error {
	s := svc.(*m.Mongo)
	query := &bson.D{}
	if info == "serverInfo" {
		query = &bson.D{{"serverStatus", 1}, {"repl", 0}, {"metrics", 0}, {"locks", 0}}
	} else {
		query = &bson.D{{info, 1}}
	}
	result := bson.M{}
	err := s.Session.DB("admin").Run(query, &result)
	if err != nil {
		return err
	}
	s.Result = result
	return nil
}

func init() {
	register("mongo", mongo)
}
