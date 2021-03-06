package ping

import (
	"log"

	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/service"
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
			svc := service.SetInit(c)
			service.DoCommand(c, svc, exec)
		},
	}
}

func init() {
	registry.RegisterCommand(command)
}
