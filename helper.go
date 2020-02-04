package dbterm

import (
	"fmt"

	"github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/setup"
	"golang.org/x/crypto/ssh/terminal"
)

func setConfigByYaml(c config.Config) error {
	v, err := setup.GetYamlConfig()
	if err != nil {
		return err
	}
	if c[config.HOST] == "" {
		c[config.HOST] = v.Driver[c[config.DRIVER].(string)].Host
	}
	if c[config.PORT] == 0 {
		c[config.PORT] = v.Driver[c[config.DRIVER].(string)].Port
	}
	if c[config.USERNAME] == "" {
		c[config.USERNAME] = v.Driver[c[config.DRIVER].(string)].Username
	}
	if c[config.PASSWORD] == "" && !c[config.FLAG_PROMPT].(bool) {
		c[config.PASSWORD] = v.Driver[c[config.DRIVER].(string)].Password
	}
	if c[config.DBNAME] == "" {
		c[config.DBNAME] = v.Driver[c[config.DRIVER].(string)].DBName
	}
	if c[config.COLLECTION] == "" {
		c[config.COLLECTION] = v.Driver[c[config.DRIVER].(string)].DBName
	}
	if v.Beautify {
		c[config.FLAG_BEAUTY] = v.Beautify
	}
	if v.Prompt {
		c[config.FLAG_PROMPT] = v.Prompt
	}
	return nil
}

func promptPassword(c config.Config) error {
	print("Input database password : ")
	passDb, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	c[config.PASSWORD] = string(passDb)
	println()
	return nil
}

func DoPrint(c config.Config, res interface{}) error {
	if c[config.FLAG_BEAUTY].(bool) {
		if err := helper.PrintPretty(res); err != nil {
			return err
		}
		return nil
	}

	fmt.Printf("%v\n", res)
	return nil
}

func FlagHelper(flg string, list map[string]string) error {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
	return nil
}

func DriverHelper(command, driver string, list []string) error {
	fmt.Printf("Error: %s not supported on selected database: %s \n\nTry using:\n", command, driver)
	for _, v := range list {
		fmt.Printf("\t%s\n", v)
	}
	println()
	return nil
}
