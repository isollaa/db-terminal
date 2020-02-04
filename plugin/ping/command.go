package ping

import (
	"fmt"
	"os"
	"strings"

	"github.com/isollaa/dbterm"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/spf13/cobra"
)

func command(parser registry.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Database availability check",
		Run: func(cmd *cobra.Command, args []string) {
			c := parser(cmd)
			t := strings.Title(strings.ToLower(c[config.CATEGORY].(string)))
			command, supported := registry.Driver(t, cmd.Use)
			if !supported {
				fmt.Printf("Error: Ping not supported on selected database: %s \n", c[config.DRIVER])
				os.Exit(1)
			}
			fmt.Printf("--%s\nPinging %s...\n", c[config.DRIVER], c[config.HOST])
			r := registry.Result{}
			if err := command(&r, c); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			dbterm.DoPrint(c, r.Value)
		},
	}
}

func init() {
	registry.RegisterCommand(command)
}
