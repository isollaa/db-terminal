package info

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/util"
)

type mongo struct {
	Result interface{}
}

func (m *mongo) Info(config dbterm.Config) error {
	session, err := util.MongoDial(config)
	if err != nil {
		return err
	}
	query := &bson.D{}
	info := fmt.Sprintf("%sInfo", config[dbterm.FLAG_STAT])
	if info == "serverInfo" {
		query = &bson.D{{"serverStatus", 1}, {"repl", 0}, {"metrics", 0}, {"locks", 0}}
	} else {
		query = &bson.D{{info, 1}}
	}
	result := bson.M{}
	err = session.DB("admin").Run(query, &result)
	if err != nil {
		return err
	}
	m.Result = result

	return nil
}

func init() {
	supportedDB["mongo"] = &mongo{}
}
