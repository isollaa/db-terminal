package dbterm

import (
	"fmt"

	"github.com/isollaa/conn/helper"
	"golang.org/x/crypto/ssh/terminal"
)

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

func Validator(flg string, list map[string]string) error {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
	return nil
}
