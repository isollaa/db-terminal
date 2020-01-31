// +build ping

package ping

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/isollaa/dbterm"
	"github.com/spf13/cobra"
)

var supportedDB = map[string]commander{}

type commander interface {
	Ping(dbterm.Config) error
}

func command(parser dbterm.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Database availability check",
		Run: func(cmd *cobra.Command, args []string) {
			config := parser(cmd)
			t := config[dbterm.CATEGORY].(string)
			command, supported := supportedDB[t]
			if !supported {
				fmt.Printf("Ping not supported on selected database: %s \n", config[dbterm.DRIVER])
				os.Exit(1)
			}
			fmt.Printf("--%s\nPinging %s...\n", config[dbterm.DRIVER], config[dbterm.HOST])
			start := time.Now()
			defer func() {
				if err := dbterm.DoPrint(config, fmt.Sprintf("Ping done in %d ms", time.Now().Sub(start).Microseconds())); err != nil {
					log.Print("unable to print: ", err)
					return
				}
			}()

			if err := command.Ping(config); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}

func init() {
	dbterm.RegisterCommand(command)
	supportedDB["mongo"] = &mongo{}
	supportedDB["sql"] = &sql{}
}
