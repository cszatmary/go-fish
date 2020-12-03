package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/TouchBistro/goutils/file"
	"github.com/cszatmary/go-fish/git"
	"github.com/cszatmary/go-fish/hooks"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var conf config

type config struct {
	SkipCI bool            `yaml:"skipCI"`
	Hooks  map[string]Hook `yaml:"hooks"`
}

type Hook struct {
	Run string `yaml:"run"`
}

func SkipCI() bool {
	return conf.SkipCI
}

func GetHook(name string) (Hook, bool) {
	hook, ok := conf.Hooks[name]
	return hook, ok
}

func Init() error {
	// Config file must be in the root dir of the repo
	rootDir, err := git.RootDir()
	if err != nil {
		return err
	}

	path := filepath.Join(rootDir, "go-fish.yml")
	log.WithFields(log.Fields{
		"file": path,
	}).Debug("Reading config file")

	if !file.FileOrDirExists(path) {
		return errors.Errorf("config file %s does not exist", path)
	}

	f, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "failed to open config file %s", path)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return errors.Wrap(err, "failed to parse config file")
	}

	for hookName, hook := range conf.Hooks {
		if !hooks.IsValidHook(hookName) {
			return errors.Errorf("%s is not a valid git hook", hookName)
		}

		if strings.TrimSpace(hook.Run) == "" {
			return errors.Errorf("hook: %s: run field cannot be empty", hookName)
		}
	}

	log.Debug("Config loaded")
	return nil
}
