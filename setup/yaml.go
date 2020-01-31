package setup

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type Yaml struct {
	Driver map[string]struct {
		Host       string `json:"host"`
		Port       int    `json:"port"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		DBName     string `json:"dbname"`
		Collection string `json:"collection"`
	} `json:"driver"`
	Beautify bool `json:"beauty"`
	Prompt   bool `json:"prompt"`
}

func GetYamlConfig() (*Yaml, error) {
	var y *Yaml
	v, err := ioutil.ReadFile("../setup/config.yaml")
	if err != nil {
		return y, err
	}

	err = yaml.Unmarshal(v, &y)
	if err != nil {
		return y, err
	}
	return y, nil
}
