package disk

import (
	"fmt"
	"log"
	"os"
	"strings"

	h "github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/helper"
	"github.com/isollaa/dbterm/registry"
	"github.com/spf13/cobra"
)

var listAttributes = map[string]string{
	config.FLAG_DB:   "disk usage of selected database",
	config.FLAG_COLL: "disk usage of selected collection",
}

func command(parser registry.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "disk",
		Short: "Database disk usage status",
		Run: func(cmd *cobra.Command, args []string) {
			c := parser(cmd)
			if err := config.RequirementCheck(c, config.FLAG_STAT); err != nil {
				log.Fatalf("Error: %s", err)
				return
			}
			if c[config.FLAG_STAT] == config.FLAG_COLL {
				if err := config.RequirementCheck(c, config.DBNAME, config.COLLECTION); err != nil {
					log.Fatalf("Error: %s", err)
					os.Exit(1)
				}
			}
			t := c[config.CATEGORY].(string)
			command, supported := registry.Driver(t, h.GetName(h.PACKAGE, command))
			if !supported {
				fmt.Printf("Error: Disk not supported on selected database: %s \n", c[config.DRIVER])
				os.Exit(1)
			}
			r := registry.Result{}
			if err := command(&r, c); err != nil {
				if strings.Contains(fmt.Sprintf("%s", err), "no such command:") {
					helper.HintFlag(c[config.FLAG_STAT].(string), listAttributes)
				}
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			helper.DoPrint(c, r.Value)
		},
	}
}

func init() {
	registry.RegisterCommand(command)
}
