package service

import (
	"log"

	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
)

func DoCommand(c t.Config, svc registry.Initial, commandFunc func(t.Config, registry.Initial) error) {
	if err := Connect(c, svc); err != nil {
		log.Print("unable to connect: ", err)
		return
	}
	defer svc.Close()
	if err := commandFunc(c, svc); err != nil {
		log.Print("error due executing command: ", err)
		return
	}
	if err := DoPrint(c, svc); err != nil {
		log.Print("unable to print: ", err)
		return
	}
}
