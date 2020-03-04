package mockgopher

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Blueprint struct {
	Host   string
	Port   int
	Routes []*Route
}

func NewBlueprint(host string, port int) *Blueprint {
	return &Blueprint{host, port, []*Route{}}
}

func (b *Blueprint) AddRoute(path string, method string, body string) *Route {
	route := &Route{path, method, body}
	b.Routes = append(b.Routes, route)
	return route
}

func (b *Blueprint) MakeRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range b.Routes {
		router.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintf(w, route.Body)
		}).Methods(route.Method)
	}
	return router
}
