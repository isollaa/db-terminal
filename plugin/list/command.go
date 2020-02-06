package list

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

var listAttributes = map[string]string{
	config.FLAG_DB:   "list databases on selected driver",
	config.FLAG_COLL: "list collection on selected database",
}

func command(parser registry.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Database attribute list",
		Run: func(cmd *cobra.Command, args []string) {
			c := parser(cmd)
			if err := config.RequirementCheck(c, config.FLAG_STAT); err != nil {
				log.Fatalf("error: %s", err)
				return
			}
			if c[config.FLAG_STAT] == config.FLAG_COLL {
				if err := config.RequirementCheck(c, config.DBNAME); err != nil {
					log.Fatalf("error: %s", err)
					os.Exit(1)
				}
			}
			t := c[config.CATEGORY].(string)
			command, supported := registry.Driver(t, h.GetName(h.PACKAGE, command))
			if !supported {
				fmt.Printf("Error: List not supported on selected database: %s \n", c[config.DRIVER])
				os.Exit(1)
			}
			r := registry.Result{}
			if err := command(&r, c); err != nil {
				helper.HintFlag(c[config.FLAG_STAT].(string), listAttributes)
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
