package info

import (
	"log"

	"github.com/isollaa/dbterm/registry"
)

type factory func(string, registry.Initial) error

var listCommand = make(map[string]factory)

func register(name string, list factory) {
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

func new(name string) factory {
	if listCommand[name] == nil {
		log.Fatalf("command for driver with '%s' category is not available.\n\nUse `app [command] --help` for more information about a command. ", name)
	}
	return listCommand[name]
}
