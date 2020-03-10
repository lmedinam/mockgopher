package main

import (
	"github.com/BurntSushi/toml"
	"github.com/medinam/mockgopher"
)

type Loader struct {
	Project string
	Locator *mockgopher.Locator
}

func NewLoader(project string, locator *mockgopher.Locator) *Loader {
	return &Loader{project, locator}
}

func (l *Loader) MakeBlueprint() (*mockgopher.Blueprint, error) {
	blueprint := mockgopher.NewBlueprint("localhost", 3000)

	if _, err := toml.Decode(l.Project, &blueprint); err != nil {
		return nil, err
	}

	for _, route := range blueprint.Routes {
		if len(route.Response.Resources) >= 1 {
			route.Response.Resources = l.Locator.LocateResources(route.Response.Resources)
		} else {
			tContent, err := l.Locator.ReadTemplate(route.Response.Template)
			if err != nil {
				return nil, err
			}
			route.Response.Template = string(tContent)
		}
	}

	return blueprint, nil
}
