// +build info

package info

import (
	"fmt"
	"log"
	"os"

	"github.com/isollaa/dbterm"
	"github.com/spf13/cobra"
)

const (
	SERVER = "server"
	BUILD  = "build"
)

var listInfo = map[string]string{
	SERVER:      "server info of selected driver",
	BUILD:       "build info of selected driver",
	dbterm.HOST: "host info of selected driver",
}

var supportedDB = map[string]commander{}

type commander interface {
	Info(dbterm.Config) error
}

func command(parser dbterm.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Database information",
		Run: func(cmd *cobra.Command, args []string) {
			config := parser(cmd)
			if err := dbterm.RequirementCheck(config, dbterm.FLAG_STAT); err != nil {
				log.Fatalf("error: %s", err)
				return
			}
			t := config[dbterm.CATEGORY].(string)
			command, supported := supportedDB[t]
			if !supported {
				fmt.Printf("Info not supported on selected database: %s \n", config[dbterm.DRIVER])
				os.Exit(1)
			}
			if err := command.Info(config); err != nil {
				dbterm.FlagHelper(config[dbterm.FLAG_STAT].(string), listInfo)
				fmt.Println(err)
				os.Exit(1)
			}
			dbterm.DoPrint(config, command)
		},
	}
}

func init() {
	dbterm.RegisterCommand(command)
	supportedDB["mongo"] = &mongo{}
}
