package mockgopher

import (
	"testing"
)

func TestLoadBlueprint(t *testing.T) {
	blueprint, _ := LoadBlueprint(`
		host = "0.0.0.0"
		port = 3000

		[[routes]]
			path = "/posts"
			method = "GET"
			body = "Hello World"
	`, new(MockTemplateReader))

	if blueprint.Host != "0.0.0.0" {
		t.Errorf("Host is not set correctly.")
	}

	if blueprint.Port != 3000 {
		t.Errorf("Port is not set correctly.")
	}

	if len(blueprint.Routes) != 1 {
		t.Errorf("Router should have 1 route, have %v route(s).", len(blueprint.Routes))
	}

	if blueprint.Routes[0].Path != "/posts" || blueprint.Routes[0].Method != "GET" {
		t.Errorf("Route Nº 0 is no set correctly: %v.", blueprint.Routes[0])
	}
}
