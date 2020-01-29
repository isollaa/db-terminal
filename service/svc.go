package service

import (
	"fmt"

	"github.com/isollaa/conn/helper"
	m "github.com/isollaa/db-terminal/cmd/init/mongo"
	s "github.com/isollaa/db-terminal/cmd/init/sql"
	t "github.com/isollaa/db-terminal/config"
	"github.com/isollaa/db-terminal/registry"
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
			fmt.Print(msg)
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
	svc.AutoFill(c)
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

func DoPrint(c t.Config, svc registry.Initial) error {
	var res interface{}
	if ss, ok := svc.(*m.Mongo); ok {
		res = ss.Result
	} else {
		ss := svc.(*s.SQL)
		res = ss.Result
	}

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
