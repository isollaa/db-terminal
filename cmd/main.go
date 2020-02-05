package main

import (
	"os"

	"github.com/isollaa/dbterm"
)

func main() {
	if err := dbterm.Exec(); err != nil {
		os.Exit(1)
	}
}
