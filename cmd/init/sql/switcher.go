package sql

import (
	"errors"
	"fmt"

	t "github.com/isollaa/db-terminal/config"
)

func GetSource(c t.Config) string {
	source := ""
	switch c[t.DB_DRIVER] {
	case "mysql":
		source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c[t.DB_USERNAME], c[t.DB_PASSWORD], c[t.DB_HOST], c[t.DB_PORT], c[t.DB_DBNAME])
	case "postgres":
		source = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", c[t.DB_HOST], c[t.DB_PORT], c[t.DB_USERNAME], c[t.DB_PASSWORD], c[t.DB_DBNAME])
	}
	return source
}

func GetQueryDB(c t.Config) string {
	query := ""
	switch c[t.DB_DRIVER] {
	case "mysql":
		query = "SHOW DATABASES"
	case "postgres":
		query = "SELECT datname FROM pg_database WHERE datistemplate = false"
	}
	return query
}

func GetQueryTable(c t.Config) string {
	query := ""
	switch c[t.DB_DRIVER] {
	case "mysql":
		query = "SHOW TABLES"
	case "postgres":
		query = "SELECT table_schema,table_name FROM information_schema.tables ORDER BY table_schema,table_name"
	}
	return query
}

func GetDiskSpace(info string, c t.Config) (map[string]string, error) {
	v := map[string]string{}
	switch c[t.DB_DRIVER] {
	case "mysql":
		switch info {
		case "db":
			// v["title"] = fmt.Sprintf("Table - %s", c[cc.DBNAME])
			// v["query"] = fmt.Sprintf("SELECT table_schema AS 'Db Name', Round( Sum( data_length + index_length ) / 1024 / 1024, 3) AS 'Db Size (MB)', Round( Sum( data_free ) / 1024 / 1024, 3 ) AS 'Free Space (MB)' FROM information_schema.tables GROUP BY table_schema")
			return v, errors.New("disk status not available")
		case "coll":
			v["title"] = fmt.Sprintf("Table - %s", c[t.DB_COLLECTION])
			v["query"] = fmt.Sprintf("SELECT (data_length+index_length)/power(1024,1) FROM information_schema.tables WHERE table_schema='%s' and table_name='%s'", c[t.DB_DBNAME], c[t.DB_COLLECTION])
		}
	case "postgres":
		switch info {
		case "db":
			info = "pg_database_size"
			v["title"] = "DB - " + c[t.DB_DBNAME].(string)
		case "coll":
			info = "pg_total_relation_size"
			c[t.DB_DBNAME] = c[t.DB_COLLECTION].(string)
			v["title"] = fmt.Sprintf("Table - %s", c[t.DB_COLLECTION])
		}
		v["query"] = fmt.Sprintf("SELECT pg_size_pretty(%s('%s'))", info, c[t.DB_DBNAME])
	default:
		return v, errors.New("disk status not available")
	}
	return v, nil
}
