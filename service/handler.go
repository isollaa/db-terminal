package service

import (
	"log"

	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
)

func DoCommand(c t.Config, commandFunc func(c t.Config, svc registry.Initial) error) {
	c[t.DB_CATEGORY] = c[t.DB_DRIVER]
	if c[t.DB_CATEGORY] == "postgres" || c[t.DB_CATEGORY] == "mysql" {
		c[t.DB_CATEGORY] = "sql"
	}
	svc := registry.NewDBInit(c[t.DB_CATEGORY].(string))
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
