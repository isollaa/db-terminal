package service

import (
	"fmt"
	"log"

	"github.com/isollaa/conn/helper"
	m "github.com/isollaa/dbterm/cmd/init/mongo"
	s "github.com/isollaa/dbterm/cmd/init/sql"
	t "github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/registry"
	"golang.org/x/crypto/ssh/terminal"
)

func RequirementCheck(c t.Config, arg ...string) error {
	for k, v := range arg {
		flag := t.RequirementCase(v)
		msg := ""
		switch c[v] {
		case "", 0:
			msg = fmt.Sprintf("Command needs flag with argument: %s `%s`\n", flag, v)
		case false:
			msg = fmt.Sprintf("Command needs flag: %s\n", flag)
		}
		if msg != "" {
			if k == len(arg)-1 {
				return fmt.Errorf(msg)
			}
			log.Print("error: ", msg)
		}
	}
	return nil
}

func Connect(c t.Config, svc registry.Initial) error {
	fmt.Printf("--%s\n", c[t.DB_DRIVER])
	if c[t.FLAG_PROMPT].(bool) {
		if err := promptPassword(c); err != nil {
			return err
		}
	}
	err := svc.Connect(c)
	if err != nil {
		return err
	}
	return nil
}

func promptPassword(c t.Config) error {
	print("Input database password : ")
	passDb, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	c[t.DB_PASSWORD] = string(passDb)
	println()
	return nil
}

func DoPrint(c t.Config,res interface{}) error {
	if c[t.FLAG_BEAUTY].(bool) {
		if err := helper.PrintPretty(res); err != nil {
			return err
		}
		return nil
	}

	fmt.Printf("%v\n", res)
	return nil
}

func Validator(flg string, list map[string]string) error {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
	return nil
}

func SetInit(c t.Config) registry.Initial {
	if err := RequirementCheck(c, t.DB_DRIVER); err != nil {
		log.Fatalf("error: %s", err)
	}
	svc := registry.NewDBInit(c[t.DB_CATEGORY].(string))
	svc.AutoFill(c)
	return svc
}
