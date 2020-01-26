package config

import (
	"fmt"
	"io"

	"github.com/cszatma/go-fish/hooks"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	SkipCI bool            `yaml:"skipCI"`
	Hooks  map[string]Hook `yaml:"hooks"`
}

type Hook struct {
	Run string `yaml:"run"`
}

func Init(r io.Reader) error {
	dec := yaml.NewDecoder(r)
	err := dec.Decode(&config)
	if err != nil {
		errors.Wrap(err, "Failed to decode config file")
	}

	for hook, _ := range config.Hooks {
		if !hooks.IsValidHook(hook) {
			return errors.New(fmt.Sprintf("%s is not a valid git hook", hook))
		}
	}

	return nil
}

func All() *Config {
	return &config
}
