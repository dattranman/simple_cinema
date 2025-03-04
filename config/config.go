package config

import (
	"fmt"
	"os"

	"github.com/dattranman/simple_cinema/model"

	yaml "gopkg.in/yaml.v2"
)

func Load(path string) (*model.Configuration, error) {
	cfg := &model.Configuration{}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file , %s", err)
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return cfg, nil
}
