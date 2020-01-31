package dbterm

import (
	"log"
	"os"
	"path/filepath"

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
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "connection host ")
	rootCmd.PersistentFlags().IntP("port", "P", 0, "connection port (default - mongo:27017 / mysql:3306 / postgres:5432)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "database username (default - mysql:root / postgres:postgres)")
	rootCmd.PersistentFlags().String("dbname", "", "connection database name (default - mongo:xsaas_ctms / mysql:mqtt)")
	rootCmd.PersistentFlags().StringP("collection", "c", "", "connection database collection name")
	rootCmd.PersistentFlags().StringP("stat", "s", "", "connection information")
	rootCmd.PersistentFlags().StringP("type", "t", "", "connection information type")
	rootCmd.PersistentFlags().BoolP("beauty", "b", false, "show pretty version of json")
	rootCmd.PersistentFlags().BoolP("prompt", "p", false, "call password prompt")

	for _, command := range listCommand {
		rootCmd.AddCommand(command(setConfig))
	}
	return rootCmd.Execute()
}

func setConfig(cmd *cobra.Command) Config {
	c := Config{
		DRIVER:      "",
		HOST:        "",
		PORT:        0,
		USERNAME:    "",
		PASSWORD:    "",
		DBNAME:      "",
		COLLECTION:  "",
		CATEGORY:    "",
		FLAG_STAT:   "",
		FLAG_TYPE:   "",
		FLAG_BEAUTY: false,
		FLAG_PROMPT: false,
	}
	c.setConfig(cmd)
	if err := RequirementCheck(c, DRIVER); err != nil {
		log.Fatalf("error: %s", err)
	}
	c[CATEGORY] = c[DRIVER]
	if c[CATEGORY] == "postgres" || c[CATEGORY] == "mysql" {
		c[CATEGORY] = "sql"
	}
	return c
}
