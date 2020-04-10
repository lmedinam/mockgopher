package main

import (
	"github.com/BurntSushi/toml"
	"github.com/medinam/mockgopher"
)

type Loader struct {
	Project string
}

func NewLoader(project string) *Loader {
	return &Loader{project}
}

func (l *Loader) MakeBlueprint() (*mockgopher.Blueprint, error) {
	blueprint := mockgopher.NewBlueprint("localhost", 3000)

	if _, err := toml.Decode(l.Project, &blueprint); err != nil {
		return nil, err
	}

	return blueprint, nil
}
