package dbterm

import (
	"fmt"

	"github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/setup"
	"golang.org/x/crypto/ssh/terminal"
)

func setConfigByYaml(c Config) error {
	v, err := setup.GetYamlConfig()
	if err != nil {
		return err
	}
	if c[HOST] == "" {
		c[HOST] = v.Driver[c[DRIVER].(string)].Host
	}
	if c[PORT] == 0 {
		c[PORT] = v.Driver[c[DRIVER].(string)].Port
	}
	if c[USERNAME] == "" {
		c[USERNAME] = v.Driver[c[DRIVER].(string)].Username
	}
	if c[PASSWORD] == "" && !c[FLAG_PROMPT].(bool) {
		c[PASSWORD] = v.Driver[c[DRIVER].(string)].Password
	}
	if c[DBNAME] == "" {
		c[DBNAME] = v.Driver[c[DRIVER].(string)].DBName
	}
	if c[COLLECTION] == "" {
		c[COLLECTION] = v.Driver[c[DRIVER].(string)].DBName
	}
	if v.Beautify {
		c[FLAG_BEAUTY] = v.Beautify
	}
	if v.Prompt {
		c[FLAG_PROMPT] = v.Prompt
	}
	return nil
}

func promptPassword(c Config) error {
	print("Input database password : ")
	passDb, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	c[PASSWORD] = string(passDb)
	println()
	return nil
}

func DoPrint(c Config, res interface{}) error {
	if c[FLAG_BEAUTY].(bool) {
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
