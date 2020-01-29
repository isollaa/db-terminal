package main

import (
	"github.com/isollaa/db-terminal/cmd"
	_ "github.com/isollaa/db-terminal/cmd/command/ping"
	_ "github.com/isollaa/db-terminal/cmd/command/stats"
	_ "github.com/isollaa/db-terminal/cmd/db/mongo"
	_ "github.com/isollaa/db-terminal/cmd/db/sql"
)

func main() {
	cmd.Exec()
}
