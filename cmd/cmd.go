package cmd

import (
	"github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	//global
	rootCmd.PersistentFlags().StringP("driver", "d", "", "connection driver name (mongo / mysql / postgres)")
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "connection host ")
	rootCmd.PersistentFlags().IntP("port", "P", 0, "connection port (default - mongo:27017 / mysql:3306 / postgres:5432)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "database username (default - mysql:root / postgres:postgres)")
	rootCmd.PersistentFlags().String("dbname", "", "connection database name (default - mongo:xsaas_ctms / mysql:mqtt)")
	rootCmd.PersistentFlags().StringP("collection", "c", "", "connection database collection name")
	rootCmd.PersistentFlags().StringP("stat", "s", "", "connection information")
	rootCmd.PersistentFlags().StringP("type", "t", "", "connection information type")
	//optional
	rootCmd.PersistentFlags().BoolP("beauty", "b", false, "show pretty version of json")
	rootCmd.PersistentFlags().BoolP("prompt", "p", false, "call password prompt")
}

func Exec() {
	c := config.SetConfig()
	for _, v := range registry.NewCommand() {
		rootCmd.AddCommand(v(c))
	}
	rootCmd.Execute()
}
