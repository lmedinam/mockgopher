package mockgopher

import (
	"github.com/gobuffalo/plush"
)

func View(template string) (string, error) {
	ctx := plush.NewContext()
	ctx.Set("faker", NewFaker())
	return plush.Render(template, ctx)
}
