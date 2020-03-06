package mockgopher

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Blueprint struct {
	Host   string
	Port   int
	Delay  *int64
	Routes []*Route
}

func NewBlueprint(host string, port int) *Blueprint {
	var delay *int64
	delay = new(int64)
	*delay = 0

	return &Blueprint{host, port, delay, []*Route{}}
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
		route := route // make a copy of the route for use in the lambda

		hpLen := len(route.Request.Headers) * 2
		hPairs := make([]string, hpLen, hpLen)

		// Convert Headers{Key, Value} to pair Headers slice (How Mux uses headers)
		for index, header := range route.Request.Headers {
			hPairs[index*2] = header.Key
			hPairs[(index*2)+1] = header.Value
		}

		router.HandleFunc(route.Request.Path, func(w http.ResponseWriter, r *http.Request) {
			if route.Response.Delay != nil {
				time.Sleep(time.Duration(*route.Response.Delay) * time.Millisecond)
			} else if b.Delay != nil {
				time.Sleep(time.Duration(*b.Delay) * time.Millisecond)
			}

			for _, header := range route.Response.Headers {
				w.Header().Set(header.Key, header.Value)
			}

			if len(route.Response.Resources) >= 1 {
				res := route.Response.Resources
				ran := rand.New(rand.NewSource(time.Now().Unix()))

				fileContent, err := ioutil.ReadFile(res[ran.Intn(len(res))])

				if err != nil {
					log.Fatal(err)
				}

				w.Write(fileContent)
			} else {
				output, err := View(route.Response.Template)
				if err != nil {
					log.Fatalln(err)
				}

				w.WriteHeader(200)

				fmt.Fprintf(w, output)
			}
		}).Methods(route.Request.Method).HeadersRegexp(hPairs...)
	}
	return router
}
