package mockgopher

import (
	"github.com/BurntSushi/toml"
)

func LoadBlueprint(project string, tReader TemplateReader) (*Blueprint, error) {
	var blueprint *Blueprint

	if _, err := toml.Decode(project, &blueprint); err != nil {
		return nil, err
	}

	for _, route := range blueprint.Routes {
		tContent, err := tReader.ReadTemplate(route.Response.Template)
		if err != nil {
			return nil, err
		}
		route.Response.Template = string(tContent)
	}

	return blueprint, nil
}
