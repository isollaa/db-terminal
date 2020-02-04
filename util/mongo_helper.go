package util

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/isollaa/dbterm/config"
)

func MongoDial(c config.Config) (*mgo.Session, error) {
	dsn := fmt.Sprintf("mongodb://%s:%d/%s",
		c[config.HOST],
		c[config.PORT],
		c[config.DBNAME],
	)

	session, err := mgo.Dial(dsn)
	if err != nil {
		return nil, err
	}

	// TODO: add session setting

	return session, nil
}

func GetMongoListSession(c config.Config, session *mgo.Session) ([]string, error) {
	switch c[config.FLAG_STAT] {
	case config.FLAG_DB:
		return session.DatabaseNames()
	case config.FLAG_COLL:
		return session.DB(c[config.DBNAME].(string)).CollectionNames()
	}
	return nil, fmt.Errorf("no such command: '%s'", c[config.FLAG_STAT])
}

func GetMongoDiskQuery(c config.Config) (interface{}, error) {
	switch c[config.FLAG_STAT] {
	case config.FLAG_DB:
		return "dbstats", nil
	case config.FLAG_COLL:
		return &bson.D{bson.DocElem{"collstats", c[config.COLLECTION].(string)}}, nil
	}
	return nil, fmt.Errorf("no such command: '%s'", c[config.FLAG_STAT])
}
