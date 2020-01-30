package dbterm

import (
	"log"

	"github.com/isollaa/conn/helper"
	"github.com/spf13/cobra"
)

var listCommand map[string]commandFactory

type ConfigParser func(*cobra.Command) Config

type commandFactory func(ConfigParser) *cobra.Command

func RegisterCommand(list commandFactory) {
	name := helper.GetPackageName(list)
	ok := false
	for k := range listCommand {
		if name != k {
			ok = true
			continue
		}
		log.Printf("Service %s already registered !", name)
	}
	if ok || len(listCommand) == 0 {
		listCommand[name] = list
	}
}
