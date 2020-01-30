package info

import (
	"fmt"
	"log"

	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/isollaa/dbterm/service"
	"github.com/spf13/cobra"
)

const (
	SERVER = "server"
	BUILD  = "build"
)

var listInfo = map[string]string{
	SERVER:    "server info of selected driver",
	BUILD:     "build info of selected driver",
	t.DB_HOST: "host info of selected driver",
}

func exec(c t.Config, svc registry.Initial) error {
	str := fmt.Sprintf("%sInfo", c[t.FLAG_STAT])
	cmd := new(c[t.DB_CATEGORY].(string))
	err := cmd(str, svc)
	if err != nil {
		service.Validator(c[t.FLAG_STAT].(string), listInfo)
		return err
	}
	return nil
}

func command(c t.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Get information of selected connection",
		Run: func(cmd *cobra.Command, args []string) {
			c.SetFlag(cmd)
			svc := service.SetInit(c)
			if err := service.RequirementCheck(c, t.FLAG_STAT); err != nil {
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
