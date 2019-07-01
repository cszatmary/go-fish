package config

import (
	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	SkipCI bool            `yaml:skipCI`
	Hooks  map[string]Hook `yaml:hooks`
}

type Hook struct {
	Run   string `yaml:run`
	Match string `yaml:match`
}

func Init(path string) error {
	config = &Config{true, nil}

	viper.AddConfigPath(path)
	viper.SetConfigName("go-fish")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)

	return err
}

func All() *Config {
	return config
}
