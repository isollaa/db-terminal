package main

import (
	"os"

	"github.com/isollaa/dbterm"
	_ "github.com/isollaa/dbterm/plugin/disk"
	_ "github.com/isollaa/dbterm/plugin/info"
	_ "github.com/isollaa/dbterm/plugin/list"
	_ "github.com/isollaa/dbterm/plugin/ping"
	_ "github.com/isollaa/dbterm/register"
	_ "github.com/isollaa/dbterm/util/sql"
)

func main() {
	if err := dbterm.Exec(); err != nil {
		os.Exit(1)
	}
}
