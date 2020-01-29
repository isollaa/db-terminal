package stats

import (
	"log"

	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
	"github.com/isollaa/db-terminal/service"
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
	valid := false
	if c[t.FLAG_STAT] == DISK {
		if err := service.RequirementCheck(c, t.FLAG_TYPE); err != nil {
			return err
		}
		if c[t.DB_DRIVER] == "postgres" {
			if err := service.RequirementCheck(c, t.FLAG_PROMPT); err != nil {
				return err
			}
		}
		if c[t.FLAG_TYPE] == t.FLAG_COLL {
			if err := service.RequirementCheck(c, t.DB_COLLECTION); err != nil {
				return err
			}
		}
		for k := range listStatusType {
			if c[t.FLAG_TYPE] == k {
				cmd := new(c[t.DB_CATEGORY].(string))
				err := cmd(c, k, svc)
				if err != nil {
					return err
				}
				valid = true
				return nil
			}
		}
		if !valid {
			service.Validator(c[t.FLAG_TYPE].(string), listStatusType)
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
			if err := service.RequirementCheck(c, t.DB_DRIVER, t.FLAG_STAT); err != nil {
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
