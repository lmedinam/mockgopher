package mockgopher

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Blueprint represent how should be serve the routes
type Blueprint struct {
	Host            string
	Port            int
	Delay           *int64
	Routes          []*Route
	Log             io.Writer
	ResourceLocator ResourceLocator
}

// NewBlueprint creates a new instance with some default values
func NewBlueprint(host string, port int) *Blueprint {
	var delay *int64
	delay = new(int64)
	*delay = 0

	return &Blueprint{host, port, delay, []*Route{}, NewStdout(), nil}
}

// AddRoute is a simple way to add new routes to a created blueprint
func (b *Blueprint) AddRoute(path string, method string, body string) *Route {
	methods := []string{method}
	route := &Route{
		Request: &Request{
			Path:    path,
			Methods: methods,
		},
		Response: &Response{
			Template: body,
			Status:   200,
		},
	}

	b.Routes = append(b.Routes, route)

	return route
}

// MakeRouter create a router using the blueprint routes
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
			var delay int64
			delay = 0

			if route.Response.Delay != nil {
				delay = *route.Response.Delay
			} else if b.Delay != nil {
				delay = *b.Delay
			}

			b.Log.Write([]byte(fmt.Sprintf(
				"%s - %s - %s - Delay: %dms\n", r.Method, r.Host, r.RequestURI, delay)))

			time.Sleep(time.Duration(delay) * time.Millisecond)

			for _, header := range route.Response.Headers {
				w.Header().Set(header.Key, header.Value)
			}

			if len(route.Response.Resources) >= 1 {
				res := route.Response.Resources
				ran := rand.New(rand.NewSource(time.Now().Unix()))

				fileContent := b.ResourceLocator.Locate(res[ran.Intn(len(res))])

				w.Write(fileContent)
			} else {
				output, err := View(string(b.ResourceLocator.Locate(route.Response.Template)))
				if err != nil {
					log.Fatalln(err)
				}

				w.WriteHeader(200)

				fmt.Fprintf(w, output)
			}
		}).Methods(route.Request.Methods...).HeadersRegexp(hPairs...)
	}
	return router
}

// Stdout is used to print stuff in the standar output implementing io.Writer
type Stdout struct {
	stdout *os.File
}

// NewStdout return a Stdout object thats use the standar output
func NewStdout() *Stdout {
	return &Stdout{os.Stdout}
}

// Write to standar output as string
func (s *Stdout) Write(p []byte) (n int, err error) {
	return fmt.Fprint(s.stdout, string(p))
}
