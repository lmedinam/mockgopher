package mockgopher

type Locator struct {
	TemplateReader
	ResourceLocater
}

func NewLocator(tReader TemplateReader, rLocater ResourceLocater) *Locator {
	return &Locator{
		tReader,
		rLocater,
	}
}
