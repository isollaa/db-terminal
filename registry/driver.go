package registry

import (
	"log"

	"github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/config"
)

type Result struct {
	Value interface{}
}

var supportedDB = make(map[string]map[string]commander)

type commander func(*Result, config.Config) error

func RegisterDriver(list commander) {
	name := helper.GetName(helper.PACKAGE, list)
	driver := helper.GetName(helper.METHOD, list)
	if supportedDB[driver][name] != nil {
		log.Printf("Feature %s for driver %s already registered !", name, driver)
		return
	}
	if supportedDB[driver] == nil {
		supportedDB[driver] = make(map[string]commander)
	}
	supportedDB[driver][name] = list
}

func Driver(driver, command string) (commander, bool) {
	_, ok := supportedDB[driver][command]
	return supportedDB[driver][command], ok
}
