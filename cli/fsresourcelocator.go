package main

import (
	"io/ioutil"
	"path/filepath"
)

type FSResourceLocator struct {
	BasePath string
}

func (r *FSResourceLocator) Locate(identifier string) []byte {
	templatePath := filepath.Join(r.BasePath, identifier)
	content, err := ioutil.ReadFile(templatePath)

	if err != nil {
		panic(err)
	}

	return content
}
