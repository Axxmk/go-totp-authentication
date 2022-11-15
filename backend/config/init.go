package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var C = new(config)

func init() {
	yml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yml, C)
	if err != nil {
		panic(err)
	}
}
