package dbterm

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const (
	DRIVER     = "driver"
	HOST       = "host"
	PORT       = "port"
	USERNAME   = "username"
	PASSWORD   = "password"
	DBNAME     = "dbname"
	COLLECTION = "collection"
	CATEGORY   = "category"

	FLAG_DB     = "db"
	FLAG_COLL   = "coll"
	FLAG_STAT   = "stat"
	FLAG_TYPE   = "type"
	FLAG_BEAUTY = "beauty"
	FLAG_PROMPT = "prompt"
)

type Config map[string]interface{}

func requirementCase(v string) string {
	switch v {
	case DRIVER:
		return "-d"
	case HOST:
		return "-H"
	case PORT:
		return "-P"
	case USERNAME:
		return "-u"
	case PASSWORD:
		return "-p"
	case DBNAME:
		return "--dbname"
	case COLLECTION:
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
	return v
}

func RequirementCheck(c Config, arg ...string) error {
	for k, v := range arg {
		if v == PASSWORD {
			err := promptPassword(c)
			if err != nil {
				return err
			}
			continue
		}
		flag := requirementCase(v)
		msg := ""
		switch c[v] {
		case "", 0:
			msg = fmt.Sprintf("Command needs flag with argument: %s `%s`\n", flag, v)
		case false:
			msg = fmt.Sprintf("Command needs flag: %s\n", flag)
		}
		if msg != "" {
			if k == len(arg)-1 {
				return fmt.Errorf(msg)
			}
			log.Print("error: ", msg)
		}
	}
	return nil
}

func (c Config) setConfig(cmd *cobra.Command) {
	for key := range c {
		if key == PASSWORD || key == CATEGORY {
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
