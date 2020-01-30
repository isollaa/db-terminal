package main

import (
	"os"

	"github.com/isollaa/dbterm"
	_ "github.com/isollaa/dbterm/cmd/command/info"
	_ "github.com/isollaa/dbterm/cmd/command/list"
	_ "github.com/isollaa/dbterm/cmd/command/ping"
	_ "github.com/isollaa/dbterm/cmd/command/stats"
	_ "github.com/isollaa/dbterm/cmd/init/mongo"
	_ "github.com/isollaa/dbterm/cmd/init/sql"
)

func main() {
	if err := dbterm.Exec(); err != nil {
		os.Exit(1)
	}
}
