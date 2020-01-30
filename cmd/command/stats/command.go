package stats

import (
	"log"

	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/service"
	"github.com/spf13/cobra"
)

const DISK = "disk"

var listStatus = map[string]string{
	DISK: "status of disk space",
}

var listStatusType = map[string]string{
	t.FLAG_DB:   "status of selected database",
	t.FLAG_COLL: "status of selected collection",
}

func exec(c t.Config, svc registry.Initial) error {
	if c[t.FLAG_STAT] == DISK {
		if err := service.RequirementCheck(c, t.FLAG_TYPE); err != nil {
			return err
		}
		if c[t.FLAG_TYPE] == t.FLAG_COLL {
			if err := service.RequirementCheck(c, t.DB_COLLECTION); err != nil {
				return err
			}
		}
		cmd := new(c[t.DB_CATEGORY].(string))
		err := cmd(c, c[t.FLAG_TYPE].(string), svc)
		if err != nil {
			service.Validator(c[t.FLAG_TYPE].(string), listStatusType)
			return err
		}
	} else {
		service.Validator(c[t.FLAG_STAT].(string), listStatus)
	}
	return nil
}

func command(c t.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get status of selected connection",
		Run: func(cmd *cobra.Command, args []string) {
			c.SetFlag(cmd)
			svc := service.SetInit(c)
			if err := service.RequirementCheck(c, t.FLAG_STAT); err != nil {
				log.Fatalf("error: %s", err)
				return
			}
			if c[t.DB_DRIVER] == "postgres" {
				if err := service.RequirementCheck(c, t.FLAG_PROMPT); err != nil {
					log.Fatalf("error: %s", err)
				}
			}
			service.DoCommand(c, svc, exec)
		},
	}
}

func init() {
	registry.RegisterCommand(command)
}
