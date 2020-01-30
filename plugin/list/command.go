package list

import (
	"log"

	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/service"
	"github.com/spf13/cobra"
)

var listAttributes = map[string]string{
	t.FLAG_DB:   "list databases on selected driver",
	t.FLAG_COLL: "list collection on selected database",
}

func exec(c t.Config, svc registry.Initial) error {
	if c[t.FLAG_STAT] == t.FLAG_COLL {
		if err := service.RequirementCheck(c, t.DB_COLLECTION); err != nil {
			return err
		}
	}
	cmd := new(c[t.DB_CATEGORY].(string))
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
			svc := service.SetInit(c)
			if err := service.RequirementCheck(c, t.DB_DBNAME, t.FLAG_STAT); err != nil {
				log.Fatalf("error: %s", err)
				return
			}
			service.DoCommand(c, svc, exec)
		},
	}
}

func init() {
	registry.RegisterCommand(command)
}
