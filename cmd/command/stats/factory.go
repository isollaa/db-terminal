package stats

import (
	"log"

	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
)

type factory func(t.Config, string, registry.Initial) error

var listCommand = make(map[string]factory)

func register(name string, list factory) {
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

func new(name string) factory {
	return listCommand[name]
}
