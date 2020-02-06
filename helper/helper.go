package helper

import (
	"fmt"
	"io/ioutil"

	"github.com/isollaa/conn/helper"
	"github.com/isollaa/dbterm/config"
	"github.com/isollaa/dbterm/setup"
	"golang.org/x/crypto/ssh/terminal"
)

func SetConfigByYaml(c config.Config) error {
	raw, err := ioutil.ReadFile("../setup/config.yaml")
	if err != nil {
		return err
	}
	v, err := setup.GetYamlConfig(raw)
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
		c[config.COLLECTION] = v.Driver[c[config.DRIVER].(string)].Collection
	}
	if v.Beautify {
		c[config.FLAG_BEAUTY] = v.Beautify
	}
	if v.Prompt {
		c[config.FLAG_PROMPT] = v.Prompt
	}
	return nil
}

func PromptPassword(c config.Config) error {
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

func HintFlag(flg string, list map[string]string) error {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
	return nil
}

func HintDriver(command, driver string, list []string) error {
	fmt.Printf("Error: %s not supported on selected database: %s \n\nTry using:\n", command, driver)
	for _, v := range list {
		fmt.Printf("\t%s\n", v)
	}
	println()
	return nil
}
