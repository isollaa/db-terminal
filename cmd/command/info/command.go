package info

import (
	"fmt"
	"log"

	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
	"github.com/isollaa/db-terminal/service"
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
	cmd := new(c[t.DB_CATEGORY].(string))
	str := fmt.Sprintf("%sInfo", c[t.FLAG_STAT])
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
