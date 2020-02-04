package dbterm

import (
	"log"
	"os"
	"path/filepath"

	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"github.com/spf13/cobra"
)

func Exec() error {
	rootCmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "DB Terminal",
		Long:  "Database monitoring tool",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	rootCmd.PersistentFlags().StringP("driver", "d", "", "connection driver name (mongo / mysql / postgres)")
	rootCmd.PersistentFlags().StringP("host", "H", "", "connection host")
	rootCmd.PersistentFlags().IntP("port", "P", 0, "connection port")
	rootCmd.PersistentFlags().StringP("username", "u", "", "database username")
	rootCmd.PersistentFlags().String("dbname", "", "connection database name")
	rootCmd.PersistentFlags().StringP("collection", "c", "", "connection database collection name")
	rootCmd.PersistentFlags().StringP("stat", "s", "", "connection information")
	rootCmd.PersistentFlags().StringP("type", "t", "", "connection information type")
	rootCmd.PersistentFlags().BoolP("beauty", "b", false, "show pretty version of json")
	rootCmd.PersistentFlags().BoolP("prompt", "p", false, "call password prompt")

	for _, command := range registry.Command() {
		rootCmd.AddCommand(command(setConfig))
	}
	return rootCmd.Execute()
}

func setConfig(cmd *cobra.Command) config.Config {
	c := config.Config{
		config.DRIVER:      "",
		config.HOST:        "",
		config.PORT:        0,
		config.USERNAME:    "",
		config.PASSWORD:    "",
		config.DBNAME:      "",
		config.COLLECTION:  "",
		config.CATEGORY:    "",
		config.FLAG_STAT:   "",
		config.FLAG_TYPE:   "",
		config.FLAG_BEAUTY: false,
		config.FLAG_PROMPT: false,
	}
	setConfig(cmd)
	if err := config.RequirementCheck(c, config.DRIVER); err != nil {
		log.Fatalf("error: %s", err)
	}
	if err := setConfigByYaml(c); err != nil {
		log.Println("unable to", err)
	}
	if c[config.FLAG_PROMPT].(bool) {
		err := promptPassword(c)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}
	c[config.CATEGORY] = c[config.DRIVER]
	if c[config.CATEGORY] == "postgres" || c[config.CATEGORY] == "mysql" {
		c[config.CATEGORY] = "sql"
	}
	return c
}
