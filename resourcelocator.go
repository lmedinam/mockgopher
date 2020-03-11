package mockgopher

type ResourceLocator interface {
	Locate(identifier string) []byte
}
