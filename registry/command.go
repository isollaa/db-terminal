package registry

import (
	"log"

	"github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/config"
	"github.com/spf13/cobra"
)

var listCommand = make(map[string]commandFactory)

type ConfigParser func(*cobra.Command) config.Config

type commandFactory func(ConfigParser) *cobra.Command

func RegisterCommand(list commandFactory) {
	name := helper.GetName(helper.PACKAGE, list)
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

func Command() map[string]commandFactory {
	return listCommand
}
