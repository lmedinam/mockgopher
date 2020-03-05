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
	route := &Route{
		Request: &Request{
			Path:   path,
			Method: method,
		},
		Response: &Response{
			Template: body,
			Status:   200,
		},
	}

	b.Routes = append(b.Routes, route)

	return route
}

func (b *Blueprint) MakeRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range b.Routes {
		hpLen := len(route.Request.Headers) * 2
		hPairs := make([]string, hpLen, hpLen)

		// Convert Headers{Key, Value} to pair Headers slice (How Mux uses headers)
		for index, header := range route.Request.Headers {
			hPairs[index*2] = header.Key
			hPairs[(index*2)+1] = header.Value
		}

		router.HandleFunc(route.Request.Path, func(w http.ResponseWriter, r *http.Request) {
			for _, header := range route.Response.Headers {
				w.Header().Set(header.Key, header.Value)
			}

			output, err := View(route.Response.Template)
			if err != nil {
				log.Fatalln(err)
			}

			w.WriteHeader(200)

			fmt.Fprintf(w, output)
		}).Methods(route.Request.Method).HeadersRegexp(hPairs...)
	}
	return router
}
