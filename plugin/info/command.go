package info

import (
	"fmt"
	"log"
	"os"

	h "github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/helper"
	"github.com/isollaa/dbterm/registry"
	"github.com/spf13/cobra"
)

const (
	SERVER = "server"
	BUILD  = "build"
)

var listInfo = map[string]string{
	SERVER:      "server info of selected driver",
	BUILD:       "build info of selected driver",
	config.HOST: "host info of selected driver",
}

func command(parser registry.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Database information",
		Run: func(cmd *cobra.Command, args []string) {
			c := parser(cmd)
			if err := config.RequirementCheck(c, config.FLAG_STAT); err != nil {
				log.Fatalf("error: %s", err)
				return
			}
			t := c[config.CATEGORY].(string)
			command, supported := registry.Driver(t, h.GetName(h.PACKAGE, command))
			if !supported {
				fmt.Printf("Error: Info not supported on selected database: %s \n", c[config.DRIVER])
				os.Exit(1)
			}
			r := registry.Result{}
			if err := command(&r, c); err != nil {
				helper.HintFlag(c[config.FLAG_STAT].(string), listInfo)
				fmt.Println(err)
				os.Exit(1)
			}
			helper.DoPrint(c, r.Value)
		},
	}
}

func init() {
	registry.RegisterCommand(command)
}
