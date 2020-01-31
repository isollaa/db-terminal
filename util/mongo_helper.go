package util

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/isollaa/dbterm"
)

func MongoDial(config dbterm.Config) (*mgo.Session, error) {
	dsn := fmt.Sprintf("mongodb://%s:%d/%s",
		config[dbterm.HOST],
		config[dbterm.PORT],
		config[dbterm.DBNAME],
	)

	session, err := mgo.Dial(dsn)
	if err != nil {
		return nil, err
	}

	// TODO: add session setting

	return session, nil
}

func GetMongoListSession(config dbterm.Config, session *mgo.Session) ([]string, error) {
	switch config[dbterm.FLAG_STAT] {
	case dbterm.FLAG_DB:
		return session.DatabaseNames()
	case dbterm.FLAG_COLL:
		return session.DB(config[dbterm.DBNAME].(string)).CollectionNames()
	}
	return nil, fmt.Errorf("no such command: '%s'", config[dbterm.FLAG_STAT])
}

func GetMongoDiskQuery(config dbterm.Config) (interface{}, error) {
	switch config[dbterm.FLAG_STAT] {
	case dbterm.FLAG_DB:
		return "dbstats", nil
	case dbterm.FLAG_COLL:
		return &bson.D{bson.DocElem{"collstats", config[dbterm.COLLECTION].(string)}}, nil
	}
	return nil, fmt.Errorf("no such command: '%s'", config[dbterm.FLAG_STAT])
}
