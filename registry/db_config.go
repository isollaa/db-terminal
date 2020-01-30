package registry

import (
	"log"

	"github.com/isollaa/conn/helper"
	t "github.com/isollaa/dbterm/config"
)

type Initial interface {
	AutoFill(t.Config)
	Connect(t.Config) error
	Close()
}

type dbInit func() Initial

var listDBInit = make(map[string]dbInit)

//auto register service by its package name
func RegisterDB(list dbInit) {
	name := helper.GetPackageName(list)
	if list == nil {
		log.Panicf("Service %s does not exist.", name)
	}
	_, registered := listDBInit[name]
	if registered {
		log.Fatalf("Service %s already registered. Ignoring.", name)
	}
	listDBInit[name] = list
}

// fill parameter using selected service package name
func NewDBInit(key string) Initial {
	run := listDBInit[key]
	if run == nil {
		log.Fatalf("database initiator '%s' not available.\n\nUse `app [command] --help` for more information about a command. ", key)
	}
	return run()
}
