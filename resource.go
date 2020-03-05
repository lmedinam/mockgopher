package mockgopher

import (
	"path/filepath"
)

type ResourceLocater interface {
	LocateResources(resources []string) []string
}

type FSResourceLocater struct {
	Path string
}

func NewFSResourceLocater(basePath string) *FSResourceLocater {
	return &FSResourceLocater{basePath}
}

func (r *FSResourceLocater) LocateResources(resources []string) []string {
	for index, res := range resources {
		resources[index] = filepath.Join(r.Path, res)
	}
	return resources
}
