package mockgopher

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gobuffalo/plush"
)

type TemplateReader interface {
	ReadTemplate(identifier string) ([]byte, error)
}

type FSTemplateReader struct {
	Path string
}

func NewFSTemplateReader(templatesPath string) *FSTemplateReader {
	return &FSTemplateReader{templatesPath}
}

func (t *FSTemplateReader) ReadTemplate(identifier string) ([]byte, error) {
	templatePath := filepath.Join(t.Path, identifier)
	content, err := ioutil.ReadFile(templatePath)

	if err != nil {
		return nil, err
	}

	return content, nil
}

type MockTemplateReader struct{}

func (t *MockTemplateReader) ReadTemplate(identifier string) ([]byte, error) {
	return []byte("Hello World!"), nil
}

func View(template string) (string, error) {
	ctx := plush.NewContext()

	ctx.Set("faker", NewFaker())

	return plush.Render(template, ctx)
}
