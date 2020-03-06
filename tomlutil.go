package mockgopher

import (
	"github.com/BurntSushi/toml"
)

func LoadBlueprint(project string, locator *Locator) (*Blueprint, error) {
	blueprint := NewBlueprint("localhost", 3000)

	if _, err := toml.Decode(project, &blueprint); err != nil {
		return nil, err
	}

	for _, route := range blueprint.Routes {
		if len(route.Response.Resources) >= 1 {
			route.Response.Resources = locator.LocateResources(route.Response.Resources)
		} else {
			tContent, err := locator.ReadTemplate(route.Response.Template)
			if err != nil {
				return nil, err
			}
			route.Response.Template = string(tContent)
		}
	}

	return blueprint, nil
}
