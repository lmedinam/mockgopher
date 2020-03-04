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

	"github.com/medinam/mockgopher"
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

	tReader := mockgopher.NewFSTemplateReader(
		filepath.Join(filepath.Dir(flag.Args()[0]), "templates"))

	blueprint, _ := mockgopher.LoadBlueprint(string(content), tReader)
	fmt.Printf("Blueprint: %v", blueprint)

	srv := &http.Server{
		Handler:      blueprint.MakeRouter(),
		Addr:         fmt.Sprintf("%s:%d", blueprint.Host, blueprint.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func usage() {
	log.Printf("Usage: %s <project-file>", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}
