package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func init() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		flag.Usage()
	}

	if len(flag.Args()) > 1 || len(flag.Args()) < 1 {
		log.Printf("%s accepts a project as input.", filepath.Base(os.Args[0]))
		flag.Usage()
	}

	content, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	loader := NewLoader(string(content))

	blueprint, _ := loader.MakeBlueprint()
	blueprint.ResourceLocator = &FSResourceLocator{filepath.Dir(flag.Args()[0])}

	srv := &http.Server{
		Handler:      blueprint.MakeRouter(),
		Addr:         fmt.Sprintf("%s:%d", blueprint.Host, blueprint.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Mock server running at: %s:%d\n\n", blueprint.Host, blueprint.Port)
	log.Fatal(srv.ListenAndServe())
}

func usage() {
	log.Printf("Usage: %s <project-file>", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}
