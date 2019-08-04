package config

import (
	"fmt"

	"github.com/cszatma/go-fish/util"
)

var config Config

type Config struct {
	SkipCI bool            `yaml:"skipCI"`
	Hooks  map[string]Hook `yaml:"hooks"`
}

type Hook struct {
	Run   string `yaml:"run"`
	Match string `yaml:"match"`
}

func Init(path string) error {
	if !util.FileExists(path) {
		return fmt.Errorf("%s does not exist", path)
	}

	err := util.ReadYaml(path, &config)
	return err
}

func All() *Config {
	return &config
}
