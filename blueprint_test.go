package mockgopher

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestNewBlueprint(t *testing.T) {
	blueprint := NewBlueprint("0.0.0.0", 3000)
	if blueprint.Host != "0.0.0.0" {
		t.Errorf("Host is not set correctly.")
	}
}

func TestRegisterRoutes(t *testing.T) {
	blueprint := NewBlueprint("0.0.0.0", 3000)
	blueprint.AddRoute("/posts", "GET", "")
	if len(blueprint.Routes) != 1 {
		t.Errorf("Router should have 1 route, have %v route(s).", len(blueprint.Routes))
	}
}

func TestMakeRoutes(t *testing.T) {
	blueprint := NewBlueprint("0.0.0.0", 3000)
	blueprint.AddRoute("/posts", "GET", "")

	router := blueprint.MakeRouter()
	postsRouteFound := false
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routePathRegexp, _ := route.GetPathRegexp()
		if routePathRegexp == "^/posts$" {
			postsRouteFound = true
		}
		return nil
	})

	if !postsRouteFound {
		t.Errorf("Route \"/posts\" not found in *mux.Router from MakeRouter().")
	}
}
