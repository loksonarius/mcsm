package yaml

import (
	"github.com/loksonarius/mcsm/pkg/config"
	"gopkg.in/yaml.v3"
)

func Marshal(i interface{}) ([]byte, error) {
	return yaml.Marshal(i)
}

func Unmarshal(dict config.ConfigDict, target interface{}) error {
	marshalled, err := yaml.Marshal(dict)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(marshalled, target); err != nil {
		return err
	}

	return nil
}
