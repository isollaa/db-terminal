package ping

import (
	"log"

	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
	"github.com/isollaa/db-terminal/service"
	"github.com/spf13/cobra"
)

func exec(c t.Config, svc registry.Initial) error {
	log.Printf("Pinging %s ", c[t.DB_HOST])
	cmd := new(c[t.DB_CATEGORY].(string))
	err := cmd(svc)
	if err != nil {
		return err
	}
	return nil
}

func command(c t.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Check ping of selected connection",
		Run: func(cmd *cobra.Command, args []string) {
			c.SetFlag(cmd)
			if err := service.RequirementCheck(c, t.DB_DRIVER); err != nil {
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
