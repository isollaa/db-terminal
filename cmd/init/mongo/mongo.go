package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
)

type Mongo struct {
	DBName     string
	Collection string
	Session    *mgo.Session
	Result     interface{}
}

func (m *Mongo) AutoFill(c t.Config) {
	if c[t.DB_PORT] == 0 {
		c[t.DB_PORT] = 27017
	}
	if c[t.DB_DBNAME] == "" {
		c[t.DB_DBNAME] = "xsaas_ctms"
	}
}

func (m *Mongo) Connect(c t.Config) error {
	source := fmt.Sprintf("%s:%d", c[t.DB_HOST], c[t.DB_PORT])
	session, err := mgo.Dial(source)
	if err != nil {
		return err
	}
	m.DBName = c[t.DB_DBNAME].(string)
	m.Collection = c[t.DB_COLLECTION].(string)
	m.Session = session
	return nil
}

func (m *Mongo) Close() {
	defer m.Session.Close()
}

func this() registry.Initial {
	return &Mongo{}
}

func init() {
	registry.RegisterDB(this)
}
