package ping

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/service"
	"github.com/spf13/cobra"
)

var supportedDB = map[string]Commander{}

type Commander interface {
	Ping(dbterm.Config) error
}

func command(parser dbterm.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Database availability check",
		Run: func(cmd *cobra.Command, args []string) {
			config := parser(cmd)

			t := config[dbterm.DB_DRIVER].(string)
			if t == "mysql" || t == "postgres" {
				t = "sql"
			}
			commander, supported := supportedDB[t]
			if !supported {
				fmt.Printf("Ping not supported on selected database: %s \n", t)
				os.Exit(1)
			}

			start := time.Now()
			defer func() {
				if err := service.DoPrint(config, fmt.Sprintf("Ping done in %d ms \n", time.Now().Sub(start).Microseconds())); err != nil {
					log.Print("unable to print: ", err)
					return
				}
			}()

			if err := commander.Ping(config); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}

func init() {
	dbterm.RegisterCommand(command)
}
