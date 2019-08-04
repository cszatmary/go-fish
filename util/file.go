package util

import (
	"os"

	"gopkg.in/yaml.v2"
)

// FileExists returns a bool indicating if a file at the given path exists.
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// ReadYaml reads a yaml file at the given path and decodes it into the given value
func ReadYaml(path string, val interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	dec := yaml.NewDecoder(file)
	err = dec.Decode(val)
	return err
}
