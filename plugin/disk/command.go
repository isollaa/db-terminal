package disk

import (
	"fmt"
	"log"
	"os"

	"github.com/isollaa/dbterm"
	"github.com/spf13/cobra"
)

var listAttributes = map[string]string{
	dbterm.FLAG_DB:   "disk usage of selected database",
	dbterm.FLAG_COLL: "disk usage of selected collection",
}

var supportedDB = map[string]commander{}

type commander interface {
	Disk(dbterm.Config) error
}

func command(parser dbterm.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "disk",
		Short: "Database disk usage status",
		Run: func(cmd *cobra.Command, args []string) {
			config := parser(cmd)
			if err := dbterm.RequirementCheck(config, dbterm.FLAG_STAT, dbterm.PASSWORD); err != nil {
				log.Fatalf("error: %s", err)
				return
			}
			if config[dbterm.FLAG_STAT] == dbterm.FLAG_COLL {
				if err := dbterm.RequirementCheck(config, dbterm.DBNAME, dbterm.COLLECTION); err != nil {
					log.Fatalf("error: %s", err)
					os.Exit(1)
				}
			}
			t := config[dbterm.CATEGORY].(string)
			command, supported := supportedDB[t]
			if !supported {
				fmt.Printf("List not supported on selected database: %s \n", config[dbterm.DRIVER])
				os.Exit(1)
			}
			if err := command.Disk(config); err != nil {
				dbterm.Validator(config[dbterm.FLAG_STAT].(string), listAttributes)
				fmt.Println(err)
				os.Exit(1)
			}
			dbterm.DoPrint(config, command)
		},
	}
}

func init() {
	dbterm.RegisterCommand(command)
}
