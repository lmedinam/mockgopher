package mockgopher

import (
	"github.com/BurntSushi/toml"
)

func LoadBlueprint(project string) (*Blueprint, error) {
	var blueprint *Blueprint
	if _, err := toml.Decode(project, &blueprint); err != nil {
		return nil, err
	}
	return blueprint, nil
}
