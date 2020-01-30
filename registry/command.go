package registry

import (
	"log"

	"github.com/isollaa/conn/helper"
	t "github.com/isollaa/dbterm/config"
	"github.com/spf13/cobra"
)

type commandFactory func(t.Config) *cobra.Command

var listCommand = make(map[string]commandFactory)

func RegisterCommand(list commandFactory) {
	name := helper.GetPackageName(list)
	ok := false
	for k, _ := range listCommand {
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

func NewCommand() map[string]commandFactory {
	return listCommand
}
