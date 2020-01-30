package mongo

import (
	"fmt"
	t "github.com/isollaa/db-terminal/config"
)

func GetListSession(c t.Config, s *Mongo) ([]string, error) {
	switch c[t.FLAG_STAT] {
	case t.FLAG_DB:
		return s.Session.DatabaseNames()
	case t.FLAG_COLL:
		return s.Session.DB(s.DBName).CollectionNames()
	}
	return nil, fmt.Errorf("no such command: '%s'", c[t.FLAG_STAT])
}
