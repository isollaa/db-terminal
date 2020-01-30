package list

import (
	"log"

	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
	"github.com/isollaa/db-terminal/service"
	"github.com/spf13/cobra"
)

var listAttributes = map[string]string{
	t.FLAG_DB:   "list databases on selected driver",
	t.FLAG_COLL: "list collection on selected database",
}

func exec(c t.Config, svc registry.Initial) error {
	cmd := new(c[t.DB_CATEGORY].(string))
	if c[t.FLAG_STAT] == t.FLAG_COLL {
		if err := service.RequirementCheck(c, t.DB_COLLECTION); err != nil {
			return err
		}
	}
	err := cmd(c, svc)
	if err != nil {
		service.Validator(c[t.FLAG_STAT].(string), listAttributes)
		return err
	}

	return nil
}

func command(c t.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list available database attributes",
		Run: func(cmd *cobra.Command, args []string) {
			c.SetFlag(cmd)
			if err := service.RequirementCheck(c, t.DB_DRIVER, t.DB_DBNAME, t.FLAG_STAT); err != nil {
				log.Print("error: ", err)
				return
			}
			service.DoCommand(c, exec)
		},
	}
}

func init() {
	registry.RegisterCommand(command)
}
