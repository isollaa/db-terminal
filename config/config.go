package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	DB_DRIVER     = "driver"
	DB_HOST       = "host"
	DB_PORT       = "port"
	DB_USERNAME   = "username"
	DB_PASSWORD   = "password"
	DB_DBNAME     = "dbname"
	DB_COLLECTION = "collection"
	DB_CATEGORY   = "category"

	FLAG_DB     = "db"
	FLAG_COLL   = "coll"
	FLAG_STAT   = "stat"
	FLAG_TYPE   = "type"
	FLAG_BEAUTY = "beauty"
	FLAG_PROMPT = "prompt"
)

type Config map[string]interface{}

func SetConfig() Config {
	return Config{
		DB_DRIVER:     "",
		DB_HOST:       "",
		DB_PORT:       0,
		DB_USERNAME:   "",
		DB_PASSWORD:   "",
		DB_DBNAME:     "",
		DB_COLLECTION: "",
		DB_CATEGORY:   "",
		FLAG_STAT:     "",
		FLAG_TYPE:     "",
		FLAG_BEAUTY:   false,
		FLAG_PROMPT:   false,
	}
}

func RequirementCase(v string) string {
	switch v {
	case DB_DRIVER:
		return "-d"
	case DB_HOST:
		return "-H"
	case DB_PORT:
		return "-P"
	case DB_USERNAME:
		return "-u"
	case DB_PASSWORD:
		return "-p"
	case DB_DBNAME:
		return "--dbname"
	case DB_COLLECTION:
		return "-c"
	case FLAG_STAT:
		return "-s"
	case FLAG_TYPE:
		return "-t"
	case FLAG_BEAUTY:
		return "-b"
	case FLAG_PROMPT:
		return "-p"
	}
	return ""
}

func (c Config) SetFlag(cmd *cobra.Command) {
	for key := range c {
		if key == DB_PASSWORD || key == DB_CATEGORY {
			continue
		}
		if v, err := cmd.Flags().GetString(key); err == nil {
			c[key] = v
			continue
		}
		if v, err := cmd.Flags().GetInt(key); err == nil {
			c[key] = v
			continue
		}
		if v, err := cmd.Flags().GetFloat64(key); err == nil {
			c[key] = v
			continue
		}
		if v, err := cmd.Flags().GetBool(key); err == nil {
			c[key] = v
			continue
		}
		fmt.Printf("flag %s doesn't exist\n", key)
	}
}
