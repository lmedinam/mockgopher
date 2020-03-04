package mockgopher

import (
	"fmt"
	"log"
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

			output, err := View(route.Body)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Fprintf(w, output)
		}).Methods(route.Method)
	}
	return router
}
